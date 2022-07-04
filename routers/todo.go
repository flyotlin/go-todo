package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/flyotlin/go-todo/models"
)

func GetTodo(c *gin.Context) {
	todos := models.GetAllTodos()
	fmt.Println(todos)

	c.JSON(http.StatusOK, todos)
}

func PostTodo(c *gin.Context) {
	title := c.DefaultPostForm("title", "First Commit")
	content := c.DefaultPostForm("content", "It's Fun, Isn't it?")

	todo := models.Todo{
		Title:   title,
		Content: content,
	}
	models.InsertNewTodo(&todo)
	fmt.Println(todo)

	c.JSON(http.StatusOK, gin.H{"response": "ok"})
}
