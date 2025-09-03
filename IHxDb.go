package HxCore

import "context"

type IHxDb interface {

	// SetDebugMode는 디버그 모드를 활성화/비활성화합니다.
	SetDebugMode(debug bool)

	Connect(userID string, password string, database string) (context.Context, error)
	/*
		// Connect는 ConnectString를 사용해 데이터베이스에 연결합니다.
		// 연결 실패 시 error를 반환합니다.
		Connect2(connectString string) (context.Context, error)
	*/

	//Open 데이터베이스 연결
	Open() (int, error)

	// Close는 데이터베이스 연결을 닫습니다.
	Close() error

	Query(query string, arg map[string]any) (int, error)

	/*
		// Query는 SELECT 문을 실행하고 그 결과를 QueryResult 인터페이스로 반환합니다.
		// '?'나 '$1' 같은 위치 기반 파라미터를 사용합니다.
		Query2(query string, args ...any) (int, error)
	*/
	nf() (int, error)
	RecordCount() (int, error)
	RecordSet() ([]any, error)

	next_record() (int, error)
	NextRecord() (int, error)

	f(colName string) (any, error)
	Field(colName string) (any, error)
	FieldByIndex(index int) (any, error)

	BeginTransaction() (context.Context, error)
	EndTransaction() error
	Commit() error
	Rollback() error

	next_id(sequenceName string) (int64, error)
	NextId(sequenceName string) (int64, error)
}
