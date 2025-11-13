package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shivashankar-dev-dot/gin-backend/models"
)

var todos = []models.Todo{
	{
		ID:        1,
		Title:     "Computers",
		Completed: false,
	},

	{
		ID:        2,
		Title:     "Social Science",
		Completed: false,
	},
}

func getTodos(c *gin.Context) {

	fmt.Println(c.Request.Method)
	c.JSON(200, gin.H{
		"todos": todos,
	})
}

func createTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	newTodo.ID = uint(len(todos) + 1)

	todos := append(todos, newTodo)

	fmt.Printf("%v", todos)
	c.JSON(http.StatusCreated, gin.H{
		"todos": todos,
	})
}

func getTodo(c *gin.Context) {
	var id, _ = strconv.Atoi(c.Param("id"))

	fmt.Println("Ckhjh")

	for _, todo := range todos {
		if int(todo.ID) == id {
			c.JSON(http.StatusOK, gin.H{
				"todo": todo,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"Not Found": "Todo not found",
	})
}

func main() {

	r := gin.Default()

	r.GET("/todos", getTodos)
	r.POST("/todos", createTodo)
	r.GET("/todos/:id", getTodo)

	r.Run(":8000")
}
