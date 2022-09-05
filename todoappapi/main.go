package main

import (
	"fmt"
	"log"

	"ibra/todoappapi/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	r := gin.Default()
	todoRepo := controllers.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the go todo app",
		})
	})

	todo := r.Group("/api")
	{
		todo.POST("/todo", todoRepo.CreateTodo)
		todo.GET("/todos", todoRepo.GetAllTodos)
		todo.GET("/todo/:id", todoRepo.GetTodo)
		todo.PUT("/todo/:id", todoRepo.UpdateTodo)
		todo.DELETE("/todo/:id", todoRepo.DeleteTodo)
	}

	r.Run("localhost:8080")
	fmt.Println("Server is running...")
}
