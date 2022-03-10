package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	schema = "%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local"
	// docker-compose.ymlに設定した環境変数を取得
	username       = os.Getenv("MYSQL_USER")
	password       = os.Getenv("MYSQL_PASSWORD")
	dbName         = os.Getenv("MYSQL_DATABASE")
	dbHost         = os.Getenv("MYSQL_HOST")
	datasourceName = fmt.Sprintf(schema, username, password, dbHost, dbName)
	Db             *sqlx.DB
)

func init() {
	log.Println("database setup")
	log.Println("dsn:", datasourceName)
	connection, err := sqlx.Open("mysql", datasourceName)

	if err != nil {
		panic("Could not connect to the database")
	}
	Db = connection
}
