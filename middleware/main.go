package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func mymiddleware1(c *gin.Context) {
	fmt.Println("========================> Middleware 1")
	c.Next()
	// Mean that  we cancel the request
	c.Abort()
}

func mymiddleware2(c *gin.Context) {
	fmt.Println("========================> Middleware 2")
	c.Next()
}

// Sharing variables
func mymiddleware3(c *gin.Context) {
	fmt.Println("========================> Middleware 3")
	c.Set("b1", "hello")
	c.Set("b2", 100)
	c.Next()
}

func mymiddleware4(c *gin.Context) {
	fmt.Println("========================> Middleware 4")
	b1 := c.MustGet("b1")
	// Can be empty
	b2, _ := c.Get("b2")

	fmt.Printf("b1 : %s . b2: %d \n", b1, b2)
	c.Next()

}

func main() {
	r := gin.Default()
	r.Use(mymiddleware1)
	r.Use(mymiddleware2)
	r.Use(mymiddleware3)
	r.Use(mymiddleware4)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Golang")
	})

	r.Run("localhost:8080")
	fmt.Println("Server is running...")
}
