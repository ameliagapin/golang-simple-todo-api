package controllers

import (
	"github.com/entirelyamelia/todo/models"
	"github.com/entirelyamelia/todo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct{}

func (ctrl Controller) CreateTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	title := c.PostForm("title")

	todo := models.Todo{Title: title, Completed: completed}
	db := utils.Database()
	db.Save(&todo)
	db.Close()
	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":     http.StatusCreated,
			"message":    "Todo item created successfully!",
			"resourceId": todo.ID,
		},
	)
}
