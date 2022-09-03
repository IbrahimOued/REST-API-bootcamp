package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Product struct {
	id    int     `json:"id" binding:required`
	Name  string  `json:"name" binding:required`
	Stock int     `json:"stock" binding:required`
	Price float32 `json:"price" binding:required`
}

func main() {
	r := gin.Default()

	// HTTP GET POST PUT DELETE
	r.GET("/", getHello)
	r.POST("/", postHello)
	r.PUT("/", putHello)
	r.DELETE("/", deleteHello)

	// Groupin endpoints
	r1 := r.Group("/api")
	{
		// HTTP GET POST PUT DELETE
		r1.GET("/", getHello)
		r1.POST("/", postHello)
		r1.PUT("/", putHello)
		r1.DELETE("/", deleteHello)
	}

	// Handling path params
	r.GET("/product/:id", getProductById)
	r.GET("/profile/:username", showProfile)
	r.GET("/compute/:num1/add/:num2", compute)

	// handling qyery request params
	// employee?firstname=ibrah&lastname=oued
	r.GET("employee", showEmployee)

	// Binding post data
	r.POST("/product", performProduct)
	r.POST("/products", performProducts)

	// Read env files
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	vala := os.Getenv("VALA")
	valb := os.Getenv("VALB")

	val_a, _ := strconv.ParseInt(vala, 10, 0)
	val_b, _ := strconv.ParseInt(valb, 10, 0)
	sum := val_a + val_b
	fmt.Printf("\n===============\n %d + %d = %d\n", val_a, val_b, sum)

	r.Run() //
	fmt.Println("Server is running...")
}

func performProduct(c *gin.Context) {
	var product Product

	if err := c.BindJSON(&product); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, product)
}

func performProducts(c *gin.Context) {
	var products []Product

	if err := c.BindJSON(&products); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, products)

}

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello REST APi - HTTP GET",
	})
}

func postHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello REST APi - HTTP POST",
	})
}

func putHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello REST APi - HTTP PUT",
	})
}

func deleteHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello REST APi - HTTP DELETE",
	})
}

func getProductById(c *gin.Context) {
	id := c.Param("id") // string
	idn, _ := strconv.ParseInt(id, 10, 0)

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":   idn,
		"name": "Product A",
	})

}

func showProfile(c *gin.Context) {
	username := c.Param("username") // string

	c.IndentedJSON(http.StatusOK, gin.H{
		"username": username,
	})
}

func compute(c *gin.Context) {
	num1, _ := strconv.ParseInt(c.Param("num1"), 10, 0)
	num2, _ := strconv.ParseInt(c.Param("num2"), 10, 0)

	sum := num1 + num2

	c.IndentedJSON(http.StatusOK, gin.H{
		"num1": num1,
		"num2": num2,
		"sum":  sum,
	})

}

func showEmployee(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "")
	lastname := c.DefaultQuery("lastname", "")

	id, _ := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 0)

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":        id,
		"firstname": firstname,
		"lastname":  lastname,
	})
}
