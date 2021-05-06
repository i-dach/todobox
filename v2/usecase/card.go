package usecase

import (
	"github.com/i-dach/todobox/v2/domain"
)

type card struct {
	*domain.Task
}

// Board user tasks controll object
type Card interface {
	Add()
	Cancel()
	Complete()
	SetAttribute()
}

func NewCard(name string) Card {
	c := &card{}
	if c.hasTask(name) {
		c.Task = domain.NewTask(name)
	}

	return c
}

// hasTask check already exist card's task.
func (c *card) hasTask(name string) bool {
	// TODO: get DB data

	return false
}

func (c *card) Add() {
	// TODO:
}

func (c *card) Cancel() {
	// TODO:

}

func (c *card) Complete() {
	// TODO:

}

func (c *card) SetAttribute() {
	// TODO:

}
