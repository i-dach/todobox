package todo

import (
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	database "github.com/sencondly/todobox/database"
)

func TestAddFunc(t *testing.T) {
	// 前準備
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO todo").WithArgs("I'm a test").WillReturnResult(sqlmock.NewResult(1, 1))
	// モックDBを使ったテスト
	if _, err = AddFunc(db, "I'm a test"); err != nil {
		t.Error(err)
	}

	// 返り値が期待通りかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestUpdateFunc(t *testing.T) {
	// 前準備
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM todo WHERE id = ?`)).
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description"}).
			AddRow("1", "I'm a test", ""))
	mock.ExpectExec("UPDATE todo").WithArgs("change title", "change desc", "1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// モックDBを使ったテスト
	data := database.Todo{
		ID:          "1",
		Title:       "change title",
		Description: "change desc",
	}

	if err = UpdateFunc(db, &data); err != nil {
		t.Error(err)
	}

	// 返り値が期待通りかを検証
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
