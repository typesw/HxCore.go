package HxCore

import (
	// "context"
	// "github.com/jmoiron/sqlx"
	// "strings"
	"fmt"
)

func CreateHxDb(provider HxDbProviderType) (IHxDb, error) {
	switch provider {
	case Oracle:
		return &HxDbOracle{}, nil // HxDbOracle 구조체 인스턴스 생성
	case MSSQL, MySQL, MariaDB, PostgreSQL, SQLite:
		return nil, fmt.Errorf("'%s' 드라이버는 아직 구현되지 않았습니다", provider)
	default:
		return nil, fmt.Errorf("알 수 없는 DB Provider: %s", provider)
	}
}

func NewHxDb(provider HxDbProviderType) (IHxDb, error) {
	return CreateHxDb(provider)
}

/*
func GetConnectionString(userId string, password string, database string) string {

	var dsn string = `user="scott" password="tiger" connectString="localhost:1521/XEPDB1"`
	return ""
}
*/

/*
func Connection(userId string, password string, database any) (*DBConfig, error) {

	safe, err := NewSafe()
	if err != nil {
		return nil, err
	}
	if err := env.Parse(safe); err != nil {
		return nil, err
	}

	timeSheet, err := NewTimeSheet()
	if err != nil {
		return nil, err
	}
	if err := env.Parse(timeSheet); err != nil {
		return nil, err
	}

	cfg := &DBConfigs{Safe: safe, TimeSheet: timeSheet}

	return cfg, nil


}
*/

func GetQueryWherString(queryWhereString ...string) string {
	var Result string = ""
	if len(queryWhereString) > 0 {
		Result = "\n"
		for _, v := range queryWhereString {
			Result += " " + Trim(v)
		}
	}
	return Result
}
