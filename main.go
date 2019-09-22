package main

import (
	"fmt"

	r "github.com/sencondly/todobox/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("API Server Started!!")

	router := gin.Default()
	if err := r.Router(router); err != nil {
		fmt.Errorf("[ERROR] %v", err)
	}

	router.Run(":8080")
}
