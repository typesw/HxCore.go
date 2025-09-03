package HxCore

type HxDbConfig struct {
	Driver   string `json:"driver" env:"DB_DRIVER" envDefault:"Oracle" description:"Database 드라이버"` // Oracle, MSSQL, MySQL, ... 등을 구분할 필드
	User     string `json:"user" env:"DB_USER" description:"User ID"`
	Password string `json:"password" env:"DB_PASSWORD" description:"패스워드"`
	Host     string `json:"host" env:"DB_HOST" envDefault:"localhost" description:"Database 서버명(주소)"`
	Port     int    `json:"port" env:"DB_PORT" description:"Database 접속 PORT"`
	DBName   string `json:"dbname" env:"DB_NAME" description:"DB Name or Service Name/ID(SID) or 추가 옵션"`
}

//#region HxDbProviderType

//#endregion HxDbProviderType

//#region HxDbConfig

//#endregion HxDbConfig
