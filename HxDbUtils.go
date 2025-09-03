package HxCore

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"
)

var pool *sql.DB // Database connection pool.

func DbPing(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
}

type HxDbProviderType string

const (
	Oracle     HxDbProviderType = "Oracle"
	MSSQL      HxDbProviderType = "MS-SQL"
	MySQL      HxDbProviderType = "MySQL"
	MariaDB    HxDbProviderType = "MariaDB"
	PostgreSQL HxDbProviderType = "PostgreSQL"
	SQLite     HxDbProviderType = "SQLite"
	OCI        HxDbProviderType = "Oracle"
)

var (
	// HxDbProviderType 상수를 키(key)로, 설명을 값(value)으로 갖는 맵을 만듭니다.
	HxDbProviderDescriptions = map[HxDbProviderType]string{
		Oracle:     "Oracle Database Server",
		MSSQL:      "Microsoft SQL Server",
		MySQL:      "MySQL Server",
		MariaDB:    "MariaDB Server",
		PostgreSQL: "PostgreSQL Server",
		SQLite:     "SQLite Database",
	}
	HxDbProviderDefaultPort = map[HxDbProviderType]int{
		Oracle:     1521,
		MSSQL:      1433,
		MySQL:      3306,
		MariaDB:    3306,
		PostgreSQL: 5432,
		SQLite:     -1,
	}
)

// HxDbProviderType 타입을 그냥 string 타입으로 변환해서 반환하면 됩니다.
func (drv HxDbProviderType) ToStringEx() string {
	return string(drv)
}

func GetHxProviderDescription(provider HxDbProviderType) string {
	// 맵에서 provider에 해당하는 설명을 찾아서 반환합니다.
	// 만약 맵에 없는 키라면, Go는 string의 제로 값인 "" (빈 문자열)을 반환합니다.
	return HxDbProviderDescriptions[provider]
}

func (drv HxDbProviderType) ToDescriptionEx() string {

	switch drv {
	case Oracle, MSSQL, MySQL, MariaDB, PostgreSQL, SQLite:
		return HxDbProviderDescriptions[drv]
	default:
		return "Unknown : " + string(drv)
	}
}

func ConnectionString(providerType HxDbProviderType, userId string, password string, database string) string {
	var connectionString string = ""
	var sbr strings.Builder
	//sbr.Write(nil)
	if !IsNullOrWhiteSpace(userId) && (Trim(userId) == "/" || Lower(Trim(userId)) == "sspi" || Lower(Trim(userId)) == "true") {
		strDbName := GetConnectionDataSourceString(providerType, database)
		switch providerType {
		case Oracle:
			sbr.WriteString(`user="` + userId + `"`)
			sbr.WriteString(` password="` + password + `"`)
			sbr.WriteString(` connectString="` + strDbName + `"`)
		}
		connectionString = sbr.String()
	}
	return connectionString
}
func GetConnectionDataSourceString(providerType HxDbProviderType, database string) string {
	Result := ""
	var defaultPort int = HxDbProviderDefaultPort[providerType]
	var strPattern string = `^([0-9a-zA-Z\.\-_]{1,})+(([:,]{1,1})([0-9]{1,5}))?([\/]{1,1}([0-9a-zA-Z\.\-_]{1,}))$`
	match := RegexpMatch(strPattern, database)
	if match != nil && len(match) > 1 {
		strDbHost := match[1]
		//strDbPortDelimiter := match[2]
		strDbPort := match[3]
		strDbName := match[4]
		if IsNullOrEmpty(strDbHost) != true && providerType == Oracle {
			if IsNullOrEmpty(strDbPort) {
				strDbPort = string(defaultPort)
			}
			Result = strDbHost + ":" + strDbPort + "/" + strDbName
		} else if IsNullOrEmpty(strDbHost) != true && providerType == MSSQL {
			if IsNullOrEmpty(strDbPort) {
				strDbHost = strDbHost + "," + strDbPort
			}
			if IsNullOrEmpty(strDbName) {

			}
			/*
				if (strDbPort.IsNullOrWhiteSpaceEx())
					strDbHost = string.Format("{0},{1}", strDbHost, strDbPort);
				thatConnStrBuilder.Add("Server", strDbHost);
				if (!strDbName.IsNullOrWhiteSpaceEx())
					thatConnStrBuilder.Add("Database", strDbName);*/
		}
	}
	return Result
}
