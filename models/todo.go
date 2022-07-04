package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

func InsertNewTodo(todo *Todo) {
	db := openDB()

	db.Create(todo)
}

func GetAllTodos() []Todo {
	db := openDB()

	var todos []Todo
	db.Find(&todos)

	return todos
}
