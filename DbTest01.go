package HxCore

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	fmt.Println("데이터베이스 테스트를 시작합니다...")
	drivers := sql.Drivers()

	fmt.Println("--- 각 요소를 한 줄씩 출력 ---")
	for index, driverName := range drivers {
		fmt.Printf("%d번째 드라이버: %s\n", index+1, driverName)
	}

	db, err := CreateHxDb(Oracle)
	if err != nil {
		log.Fatalf("DB 핸들러 생성 실패: %v", err)
	}
	// 프로그램 종료 시 DB 연결 자동 해제
	defer db.Close()
}
