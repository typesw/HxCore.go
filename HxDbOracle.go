package HxCore

import (
	"context"
	//"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/godror/godror" // Oracle 드라이버 (블랭크 임포트)
	"github.com/jmoiron/sqlx"
)

/**
 * HxCore 참조 파일
 * - HxDbConfig
 * - HxDbUtils
 */

// HxDbOracle는 IHxDb 인터페이스를 구현하는 Oracle 전용 구조체입니다.
type HxDbOracle struct {
	db          *sqlx.DB         // DB 커넥션 풀
	tx          *sqlx.Tx         // 트랜잭션 객체
	lastRecords []map[string]any // 마지막 쿼리 결과를 저장하는 슬라이스
	cursor      int              // 현재 레코드 커서 위치 (-1은 시작 전)
	isDebug     bool
}

func (h *HxDbOracle) SetDebugMode(debug bool) {
	h.isDebug = debug
}

// Connect는 사용자 정보로 DSN을 만들어 Oracle DB에 연결합니다.
func (h *HxDbOracle) Connect(userID string, password string, database string) (context.Context, error) {
	if h.db != nil {
		if h.isDebug {
			fmt.Println("이미 DB에 연결되어 있습니다.")
		}
		return context.Background(), nil
	}
	//connectString := GetConnectionDataSourceString(Oracle, database)
	// Oracle DSN(Data Source Name) 생성
	dsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s"`, userID, password, database)

	db, err := sqlx.Connect("godror", dsn)
	if err != nil {
		return nil, fmt.Errorf("Oracle 연결 실패: %w", err)
	}

	h.db = db
	h.cursor = -1 // 커서 초기화
	if h.isDebug {
		fmt.Println("Oracle DB에 성공적으로 연결되었습니다.")
	}

	return context.Background(), nil
}

// Open은 Connect와 동일한 역할을 수행 (인터페이스 호환성을 위해)
func (h *HxDbOracle) Open() (int, error) {
	// 실제 Open 로직은 Connect에서 처리되므로 여기서는 상태만 확인
	if h.db == nil {
		return 0, fmt.Errorf("DB가 연결되지 않았습니다. Connect를 먼저 호출하세요")
	}

	//h.db.Query("PRAGMA foreign_keys = ON")
	if h.isDebug {
		//h.db.Query("PRAGMA foreign_keys = OFF")
	}

	return 1, nil
}

// Close는 DB 연결을 닫습니다.
func (h *HxDbOracle) Close() error {
	if h.db != nil {
		if h.isDebug {
			log.Println("Oracle DB 연결을 닫습니다.")
		}
		return h.db.Close()
	}
	return nil
}

// Query는 SELECT 문을 실행하고 그 결과를 내부 레코드셋에 저장합니다.
func (h *HxDbOracle) Query(query string, arg map[string]any) (int, error) {
	if h.db == nil {
		return 0, fmt.Errorf("DB가 연결되지 않았습니다")
	}

	var rows *sqlx.Rows
	var err error

	// 트랜잭션 중이면 트랜잭션 객체로 쿼리 실행
	if h.tx != nil {
		rows, err = h.tx.NamedQuery(query, arg)
	} else {
		rows, err = h.db.NamedQuery(query, arg)
	}

	if err != nil {
		return 0, fmt.Errorf("쿼리 실행 실패: %w", err)
	}
	defer rows.Close()

	// 모든 결과를 읽어서 내부 레코드셋에 저장
	records := []map[string]any{}
	for rows.Next() {
		record := make(map[string]any)
		if err := rows.MapScan(record); err != nil {
			return 0, fmt.Errorf("결과 스캔 실패: %w", err)
		}
		// Oracle 드라이버는 모든 컬럼 이름을 대문자로 반환하는 경향이 있음
		// 사용 편의를 위해 소문자로 변환
		lowerRecord := make(map[string]any)
		for k, v := range record {
			lowerRecord[strings.ToLower(k)] = v
		}
		records = append(records, lowerRecord)
	}

	h.lastRecords = records
	h.cursor = -1 // 커서 초기화
	if h.isDebug {
		log.Printf("쿼리 실행 완료. %d개의 레코드를 가져왔습니다.", len(h.lastRecords))
	}

	return len(h.lastRecords), nil
}

// nf (number of fields/rows)는 마지막 쿼리의 레코드 수를 반환합니다.
func (h *HxDbOracle) nf() (int, error) {
	if h.lastRecords == nil {
		return 0, fmt.Errorf("먼저 Query를 실행해야 합니다")
	}
	return len(h.lastRecords), nil
}

// RecordCount는 nf와 동일한 기능을 합니다.
func (h *HxDbOracle) RecordCount() (int, error) {
	return h.nf()
}

// RecordSet은 저장된 전체 레코드셋을 반환합니다.
func (h *HxDbOracle) RecordSet() ([]any, error) {
	if h.lastRecords == nil {
		return nil, fmt.Errorf("먼저 Query를 실행해야 합니다")
	}
	// []map[string]any를 []any로 변환
	resultSet := make([]any, len(h.lastRecords))
	for i, v := range h.lastRecords {
		resultSet[i] = v
	}
	return resultSet, nil
}

// next_record는 내부 커서를 다음 레코드로 이동시킵니다.
func (h *HxDbOracle) next_record() (int, error) {
	if h.lastRecords == nil {
		return 0, fmt.Errorf("먼저 Query를 실행해야 합니다")
	}
	if h.cursor+1 < len(h.lastRecords) {
		h.cursor++
		return 1, nil // 성공
	}
	return 0, nil // 더 이상 레코드가 없음 (EOF)
}

// NextRecord는 next_record와 동일한 기능을 합니다.
func (h *HxDbOracle) NextRecord() (int, error) {
	return h.next_record()
}

// f (field)는 현재 커서 위치의 레코드에서 특정 컬럼 값을 가져옵니다.
func (h *HxDbOracle) f(colName string) (any, error) {
	if h.cursor < 0 || h.cursor >= len(h.lastRecords) {
		return nil, fmt.Errorf("유효한 레코드 위치가 아닙니다. next_record를 먼저 호출하세요")
	}
	currentRecord := h.lastRecords[h.cursor]
	value, ok := currentRecord[strings.ToLower(colName)]
	if !ok {
		return nil, fmt.Errorf("컬럼 '%s'을(를) 찾을 수 없습니다", colName)
	}
	return value, nil
}

// Field는 f와 동일한 기능을 합니다.
func (h *HxDbOracle) Field(colName string) (any, error) {
	return h.f(colName)
}

// FieldByIndex는 현재 레코드에서 인덱스로 컬럼 값을 가져옵니다. (Map에서는 순서가 보장되지 않으므로 주의)
func (h *HxDbOracle) FieldByIndex(index int) (any, error) {
	if h.cursor < 0 || h.cursor >= len(h.lastRecords) {
		return nil, fmt.Errorf("유효한 레코드 위치가 아닙니다. next_record를 먼저 호출하세요")
	}
	currentRecord := h.lastRecords[h.cursor]
	if index >= len(currentRecord) {
		return nil, fmt.Errorf("인덱스 %d가 범위를 벗어났습니다", index)
	}
	i := 0
	for _, v := range currentRecord {
		if i == index {
			return v, nil
		}
		i++
	}
	return nil, fmt.Errorf("값을 찾을 수 없습니다") // 이론적으로 도달 불가
}

// 트랜잭션 관련 메서드 (구현)
func (h *HxDbOracle) BeginTransaction() (context.Context, error) {
	if h.tx != nil {
		return nil, fmt.Errorf("이미 트랜잭션이 진행 중입니다")
	}
	tx, err := h.db.Beginx()
	if err != nil {
		return nil, err
	}
	h.tx = tx
	if h.isDebug {
		log.Println("트랜잭션을 시작합니다.")
	}
	return context.Background(), nil
}

func (h *HxDbOracle) EndTransaction() error {
	// Go에서는 Commit 또는 Rollback으로 트랜잭션을 명시적으로 종료해야 합니다.
	// 이 메서드는 두 경우를 모두 처리하기 위해 남겨둘 수 있습니다.
	if h.tx == nil {
		if h.isDebug {
			log.Println("진행 중인 트랜잭션이 없어 EndTransaction을 건너뜁니다.")
		}
		return nil
	}
	// 기본적으로 롤백하여 안전하게 처리
	return h.Rollback()
}

func (h *HxDbOracle) Commit() error {
	if h.tx == nil {
		return fmt.Errorf("시작된 트랜잭션이 없습니다")
	}
	err := h.tx.Commit()
	h.tx = nil // 트랜잭션 객체 초기화
	if err == nil && h.isDebug {
		log.Println("트랜잭션이 커밋되었습니다.")
	}
	return err
}

func (h *HxDbOracle) Rollback() error {
	if h.tx == nil {
		return fmt.Errorf("시작된 트랜잭션이 없습니다")
	}
	err := h.tx.Rollback()
	h.tx = nil // 트랜잭션 객체 초기화
	if err == nil && h.isDebug {
		log.Println("트랜잭션이 롤백되었습니다.")
	}
	return err
}

// next_id는 Oracle 시퀀스 객체에서 다음 값을 가져옵니다.
func (h *HxDbOracle) next_id(sequenceName string) (int64, error) {
	if h.db == nil {
		return 0, fmt.Errorf("DB가 연결되지 않았습니다")
	}

	// 시퀀스 이름은 파라미터로 바인딩할 수 없으므로, Sprintf를 사용합니다.
	// 이로 인해 sequenceName은 외부 사용자 입력이 아닌, 코드 내부에서 제어되는 값이어야 합니다.
	query := fmt.Sprintf("SELECT %s.NEXTVAL FROM dual", sequenceName)

	var nextId int64
	var err error

	ctx := context.Background()

	// 트랜잭션이 진행 중인 경우와 아닌 경우를 구분하여 실행합니다.
	if h.tx != nil {
		err = h.tx.QueryRowxContext(ctx, query).Scan(&nextId)
	} else {
		err = h.db.QueryRowxContext(ctx, query).Scan(&nextId)
	}

	if err != nil {
		return 0, fmt.Errorf("시퀀스 '%s' 값 가져오기 실패: %w", sequenceName, err)
	}

	if h.isDebug {
		log.Printf("시퀀스 '%s'의 다음 값: %d", sequenceName, nextId)
	}

	return nextId, nil
}

// next_id는 NextId 메서드를 호출하는 별칭(alias)입니다.
func (h *HxDbOracle) NextId(sequenceName string) (int64, error) {
	return h.next_id(sequenceName)
}
