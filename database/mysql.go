package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB_NAME = os.Getenv("DBN")
	DB_USER = os.Getenv("DB_USER")
	DB_PWD  = os.Getenv("DB_PWD")
)

// Todo = this tables into Todo Resource infomation.
type Todo struct {
	ID          string `id:"id"`
	Title       string `title:"task"`
	Description string `description:""`
}

/*********************************************
*                 DB Config                  *
*********************************************/
// Open = mysqlに接続するための関数（設定値は共通で管理）
func Open() *sql.DB {
	db, err := sql.Open("mysql", DB_USER+":"+DB_PWD+"@/"+DB_NAME)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
