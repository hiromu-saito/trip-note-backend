package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	schema = "%s:%s@tcp(mysql:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	// docker-compose.ymlに設定した環境変数を取得
	username       = os.Getenv("MYSQL_USER")
	password       = os.Getenv("MYSQL_PASSWORD")
	dbName         = os.Getenv("MYSQL_DATABASE")
	datasourceName = fmt.Sprintf(schema, username, password, dbName)
	Db             *sqlx.DB
)

func Connect() {
	// GORMセット
	fmt.Println(datasourceName)
	connection, err := sqlx.Open("mysql", datasourceName)
	if err != nil {
		panic("Could not connect to the database")
	}
	Db = connection
}
