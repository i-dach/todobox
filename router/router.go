package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sencondly/todobox/todo"
)

// GinRouter is router that API method router.
func Router(r *gin.Engine) error {

	// helth check
	r.GET("/helth", func(c *gin.Context) {
		// apm.TraceSeg(c, "/helth")

		c.JSON(200, gin.H{
			"message": "helth check ok",
		})
	})

	// Simple group
	api := r.Group("/todo")
	{
		// Entity の操作
		api.POST("/task", todo.Add)
		api.PATCH("/task", todo.Update)
		api.GET("/task", todo.Select)

		// Task event
		api.POST("/task/done/:id", todo.Done)
		api.DELETE("/task/delete", todo.Delete)
	}

	return nil
}
