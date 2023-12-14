package dbo

type ENGINE string
type DRIVER string

const (
	DRIVER_MYSQL  DRIVER = "mysql"
	DRIVER_PGSQL  DRIVER = "pgsql"
	DRIVER_SQLITE DRIVER = "sqlite"
)

const (
	ENGINE_INNODB ENGINE = "InnoDB"
)

type Options struct {
	Driver    DRIVER
	Host      string
	Port      string
	Username  string
	Password  string
	DBName    string
	Charset   string
	Collation string
	DSN       string
	Engine    ENGINE
}
