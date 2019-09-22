package todo

import "github.com/gin-gonic/gin"

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

func Add(c *gin.Context) {
	msg := "test"
	todo := Success()
	todo.Resp(c, msg)
}
