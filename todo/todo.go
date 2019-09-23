package todo

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	database "github.com/sencondly/todobox/database"
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

	db, err := database.Open()
	if err != nil {
		msg := fmt.Sprintf("db connection error: %v", err)
		todo := Error()
		todo.Resp(c, msg)
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

// Update = taskの更新を行う
/*
	更新する際にはIDの指定を必須とする
	更新可能なのは「タイトル」「内容（詳細）」
*/
func Update(c *gin.Context) {
	// patchされたデータを取得
	var data database.Todo
	if err := c.ShouldBind(&data); err != nil {
		msg := fmt.Sprintf("no setting data: %v", err)
		todo := Error()
		todo.Resp(c, msg)
		return
	}

	db, err := database.Open()
	if err != nil {
		msg := fmt.Sprintf("db connection error: %v", err)
		todo := Error()
		todo.Resp(c, msg)
		return
	}

	if err := UpdateFunc(db, &data); err != nil {
		todo := NotFound()
		todo.Resp(c, fmt.Sprint(err))
		return
	}

	msg := fmt.Sprintf("update recode: %s", data.ID)
	todo := Success()
	todo.Resp(c, msg)
}

func UpdateFunc(db *sql.DB, data *database.Todo) error {
	// トランザクションを開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow("SELECT * FROM todo WHERE id = ?", data.ID)

	var res database.Todo
	if err = row.Scan(&res.ID, &res.Title, &res.Description); err != nil {
		tx.Rollback()
		return err
	}

	const taskUpdateSQL = "UPDATE todo SET title = ?, description = ? WHERE id = ?"
	_, err = tx.Exec(taskUpdateSQL, data.Title, data.Description, data.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// Select = taskの取得を行う
/*
	Delしていないtaskを取得する
*/
func Select(c *gin.Context) {
	// 基本的にはDel以外は全取得
	var data []database.Todo

	db, err := database.Open()
	if err != nil {
		msg := fmt.Sprintf("db connection error: %v", err)
		todo := Error()
		todo.Resp(c, msg)
		return
	}

	if err := SelectFunc(db, &data); err != nil {
		todo := NotFound()
		todo.Resp(c, fmt.Sprint(err))
		return
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		msg := fmt.Sprintf("json marshal error: %v", err)
		todo := NotFound()
		todo.Resp(c, msg)
		return
	}

	msg := fmt.Sprint(jsonBytes)
	todo := Success()
	todo.Resp(c, msg)
}

func SelectFunc(db *sql.DB, d *[]database.Todo) error {
	rows, err := db.Query("SELECT id, title, description FROM todo WHERE id NOT IS (SELECT id FROM events WHERE del = 0)")
	if err != nil {
		return err
	}

	var data []database.Todo
	for rows.Next() {
		var res database.Todo
		if err := rows.Scan(&res.ID, &res.Title, &res.Description); err != nil {
			continue
		}
		data = append(data, res)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	d = &data
	return nil
}

// Done = taskを完了させる
/*
	完了はeventsテーブルに登録をする形
*/
func ç(c *gin.Context) {
	// 指定されたidを取得
	id := c.PostForm("id")
	if id == "" {
		todo := Error()
		todo.Resp(c, "no setting id")
		return
	}

	db, err := database.Open()
	if err != nil {
		msg := fmt.Sprintf("db connection error: %v", err)
		todo := Error()
		todo.Resp(c, msg)
		return
	}

	res, err := DoneFunc(db, id)
	if err != nil {
		todo := NotFound()
		todo.Resp(c, "insert record error:")
		return
	}

	msg := fmt.Sprintf("insert recode :%d", res)
	todo := Success()
	todo.Resp(c, msg)
}

func DoneFunc(db *sql.DB, id string) (int64, error) {
	const sql = "INSERT INTO events(id, done) values(?, 1)"
	r, err := db.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	return r.LastInsertId()
}
