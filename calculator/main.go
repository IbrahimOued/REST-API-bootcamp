package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Numeric struct {
	Num1   float32 `json:"num1"`
	Num2   float32 `json:"num2"`
	Result float32 `json:"sum"`
}

func main() {
	r := gin.Default()

	r.POST("/add", add)
	r.POST("/substract", substract)
	r.POST("/multiply", multiply)
	r.POST("/divide", divide)

	r.Run("localhost:8080")
	fmt.Println("server is running...")
}

func add(c *gin.Context) {
	var num Numeric
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return

	}
	num.Result = num.Num1 + num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func substract(c *gin.Context) {
	var num Numeric
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return

	}
	num.Result = num.Num1 - num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func multiply(c *gin.Context) {
	var num Numeric
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return

	}
	num.Result = num.Num1 * num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func divide(c *gin.Context) {
	var num Numeric
	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return

	}
	num.Result = num.Num1 * num.Num2
	c.IndentedJSON(http.StatusOK, num)
}
