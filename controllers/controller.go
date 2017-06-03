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

	if db.Error != nil {
		ctrl.handleException(c, &db.Error)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"status":     http.StatusCreated,
			"message":    "Todo item created successfully!",
			"resourceId": todo.ID,
		},
	)
}

func (ctrl Controller) FetchAllTodo(c *gin.Context) {
	var todos []models.Todo
	var _todos []models.TransformedTodo

	db := utils.Database()
	db.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No todo found!",
			},
		)
		return
	}

	//transforms the todos for the response
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		}

		_todos = append(_todos, models.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _todos,
		},
	)
}

func (ctrl Controller) FetchSingleTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")

	db := utils.Database()
	db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No todo found!",
			},
		)
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	}

	_todo := models.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(
		http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   _todo,
		},
	)
}

func (ctrl Controller) UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	db := utils.Database()
	db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No todo found!",
			},
		)
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	db.Model(&todo).Update("completed", c.PostForm("completed"))
	db.Close()

	if db.Error != nil {
		ctrl.handleException(c, &db.Error)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Todo updated successfully!",
		},
	)
}

func (ctrl Controller) DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Param("id")
	db := utils.Database()
	db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  http.StatusNotFound,
				"message": "No todo found!",
			},
		)
		return
	}

	db.Delete(&todo)

	if db.Error != nil {
		ctrl.handleException(c, &db.Error)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  http.StatusOK,
			"message": "Todo deleted successfully!",
		},
	)
}

func (ctrl Controller) handleException(c *gin.Context, err *error) {
	c.JSON(
		http.StatusExpectationFailed,
		gin.H{
			"status":  http.StatusExpectationFailed,
			"message": err,
		},
	)
}
