package todo

import "github.com/gin-gonic/gin"

func Add(c *gin.Context) {
	c.JSON(200, gin.H{
		"todo": "msg",
	})
}