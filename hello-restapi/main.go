package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello REST API")
	})

	r.GET("/hello2", hello)

	r.Run() // runs on port 8080 by default or r.Run("8090")
	fmt.Println("Server is running")
}

func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello REST API")
}
