package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Task      string    `json:"task"`
	Completed bool      `json:"completed"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
}

// Crud functionnalities
func CreateTodo(db *gorm.DB, newTodo *Todo) (error error) {
	error = db.Create(newTodo).Error
	if error != nil {
		return error
	}
	return nil
}

func GetTodos(db *gorm.DB, todos *[]Todo) (error error) {
	error = db.Find(todos).Error
	if error != nil {
		return error
	}
	return nil
}

func GetTodoById(db *gorm.DB, todo *Todo, id int) (error error) {
	error = db.Where("id = ?", id).First(todo).Error
	if error != nil {
		return error
	}
	return nil
}

func UpdateTodo(db *gorm.DB, todo *Todo) (error error) {
	db.Save(todo)
	return nil

}

func DeleteTodoById(db *gorm.DB, todo *Todo, id int) (error error) {
	db.Where("id = ?").Delete(todo)
	return nil
}
