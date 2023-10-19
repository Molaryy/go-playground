package todos

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"ID"`
	Item      string `json:"Item"`
	Completed bool   `json:"Completed"`
}

var todos = []todo{
	{ID: "1", Item: "How to TIE YOUR SHOELACES | Step by Step Guide for Kids", Completed: false},
	{ID: "2", Item: "Epitech > 42", Completed: true},
	{ID: "3", Item: "Buy some tomatoes", Completed: false},
}

// get todo
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

// get todo/:id

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getContextId(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint("this id was not found ", id)})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

// post todo

func addNewTodo(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// patch todo

func patchTodoById(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprint("this id was not found ", id)})
		return
	}
	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

// todo handler

func TodoHandler(router *gin.Engine) {
	router.POST("/todos", addNewTodo)
	router.GET("todos/:id", getContextId)
	router.GET("/todos", getTodos)
	router.PATCH("todos/:id", patchTodoById)
}
