// package usecase

// import (
// 	"github.com/i-dach/todobox/v2/domain"
// 	"github.com/i-dach/todobox/v2/domain/repository"
// )

// type board struct {
// 	cards []Card
// 	repo  repository.Task
// }

// // Card add user tasks with sprint
// type Board interface {
// 	Add()
// 	Cancel()
// 	Complete()
// 	// SetAttribute()
// }

// func NewBoard(repo repository.Task) Card {
// 	return &board{
// 		repo: repo,
// 	}
// }

// func (c *card) SetTask(task *domain.Task) {
// 	if task != nil {
// 		c.Task = task
// 	}

// 	return
// }
