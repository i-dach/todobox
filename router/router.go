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
	api := r.Group("/api")
	{
		api.GET("/todo", todo.Add)
		// q.GET("/:tag", todo.)
	}

	return nil
}
