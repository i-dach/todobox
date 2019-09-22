package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB_NAME = os.Getenv("DBN")
	DB_USER = os.Getenv("DB_USER")
	DB_PWD  = os.Getenv("DB_PWD")
)

/*********************************************
*                 DB Config                  *
*********************************************/
type DB interface {
	Open() (*sql.DB, error)
}

type TestDB struct {
	DB DB
}

// Open = mysqlに接続するための関数（設定値は共通で管理）
func Open() (*sql.DB, error) {
	return sql.Open("mysql", DB_USER+":"+DB_PWD+"@/"+DB_NAME)
}

/*********************************************
*             table definition               *
*********************************************/
// Todo = this tables into Todo Resource infomation.
type Todo struct {
	ID          string `id:"id"`
	Title       string `title:"task"`
	Description string `description:""`
}
