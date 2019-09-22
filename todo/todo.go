package todo

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/sencondly/todobox/database"
)

type Todo interface {
	Resp(c *gin.Context, msg string)
}

type TodoFunc func(c *gin.Context, msg string)

func (todo TodoFunc) Resp(c *gin.Context, msg string) {
	todo(c, msg)
}

func Success() Todo {
	return TodoFunc(func(c *gin.Context, msg string) {
		c.JSON(200, gin.H{
			"todo": msg,
		})
	})
}

func Error() Todo {
	return TodoFunc(func(c *gin.Context, msg string) {
		c.JSON(500, gin.H{
			"todo": msg,
		})
	})
}

func NotFound() Todo {
	return TodoFunc(func(c *gin.Context, msg string) {
		c.JSON(400, gin.H{
			"todo": msg,
		})
	})
}

// Add = taskの登録を行う
/*
	必須項目は「タイトル」だけ
	内容(詳細)の登録はupdateオンリー
*/
func Add(c *gin.Context) {
	// postされたデータを取得
	title := c.PostForm("title")
	if title == "" {
		todo := Error()
		todo.Resp(c, "no posting title")
		return
	}

	db, err := db.Open()
	if err != nil {
		todo := Error()
		todo.Resp(c, "db connection error")
		return
	}

	id, err := AddFunc(db, title)
	if err != nil {
		todo := NotFound()
		todo.Resp(c, "insert record error:")
		return
	}

	msg := fmt.Sprintf("insert recode :%d", id)
	todo := Success()
	todo.Resp(c, msg)
}

func AddFunc(db *sql.DB, title string) (int64, error) {
	const sql = "INSERT INTO todo(title) values(?)"
	r, err := db.Exec(sql, title)
	if err != nil {
		return 0, err
	}

	return r.LastInsertId()
}
