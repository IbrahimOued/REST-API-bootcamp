package main

import (
	"fmt"
	"ibra/productapi/controllers"
	"ibra/productapi/models"

	"github.com/gin-gonic/gin"
)

var products = []models.Product{
	{Id: 1, Name: "Product 1", Stock: 10, Price: 2.5},
	{Id: 2, Name: "Product 2", Stock: 19, Price: 13.5},
	{Id: 3, Name: "Product 3", Stock: 40, Price: 3.5},
}

func main() {
	r := gin.Default()
	controller := controllers.Init(&products)

	r.GET("/products", controller.ReadProducts)
	r.GET("/product/:id", controller.ReadProductById)
	r.POST("/product", controller.CreateProduct)
	r.PUT("/product/:id", controller.UpdateProductById)
	r.DELETE("/product/:id", controller.DeleteProductById)

	r.Run("localhost:8080")
	fmt.Println("Server is running")
}
