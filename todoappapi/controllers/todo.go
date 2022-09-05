package controllers

import (
	"errors"
	"ibra/todoappapi/database"
	"ibra/todoappapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Dependency injection
type TodoRepo struct {
	Db *gorm.DB
}

func New() *TodoRepo {
	db := database.InitDB()
	db.AutoMigrate(&models.Todo{})
	return &TodoRepo{Db: db}
}

func (repository *TodoRepo) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if c.BindJSON(&todo) == nil {
		error := models.CreateTodo(repository.Db, &todo)
		if error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": error})
			return
		}
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusBadRequest, todo)
	}

}

func (repository *TodoRepo) GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	error := models.GetTodos(repository.Db, &todos)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": error,
		})
	}
	c.JSON(http.StatusOK, todos)
}

func (repository *TodoRepo) GetTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	idn, _ := strconv.Atoi(id)
	var todo models.Todo

	error := models.GetTodoById(repository.Db, &todo, idn)
	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": error})
	}
	c.JSON(http.StatusOK, todo)

}

func (repository *TodoRepo) UpdateTodo(c *gin.Context) {
	var todo models.Todo
	var updatedTodo models.Todo

	id, _ := c.Params.Get("id")
	idn, _ := strconv.Atoi(id)
	err := models.GetTodoById(repository.Db, &updatedTodo, idn)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if c.BindJSON(&todo) == nil {
		updatedTodo.Task = todo.Task
		updatedTodo.Completed = todo.Completed
		updatedTodo.StartDate = todo.StartDate
		updatedTodo.EndDate = todo.EndDate

		error := models.UpdateTodo(repository.Db, &updatedTodo)
		if error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"eroor": error})
			return
		}
		c.JSON(http.StatusOK, updatedTodo)
	} else {
		c.JSON(http.StatusBadRequest, todo)
	}

}

func (repository *TodoRepo) DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id, _ := c.Params.Get("id")
	idn, _ := strconv.Atoi(id)

	error := models.DeleteTodoById(repository.Db, &todo, idn)
	if error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo was deleted successfully"})

}
