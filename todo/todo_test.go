package todo

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestAdd(t *testing.T) {
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
