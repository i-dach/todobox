package repository

import (
	"database/sql"

	"github.com/i-dach/todobox/v2/domain"
)

type Task interface {
	Add(DB *sql.DB, task *domain.Task) error
	List(DB *sql.DB) ([]*domain.Task, error)
	Update(DB *sql.DB, task *domain.Task) error
}
