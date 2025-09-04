package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("데이터베이스 테스트를 시작합니다...")
	drivers := sql.Drivers()

	fmt.Println("--- 각 요소를 한 줄씩 출력 ---")
	for index, driverName := range drivers {
		fmt.Printf("%d번째 드라이버: %s\n", index+1, driverName)
	}
}
