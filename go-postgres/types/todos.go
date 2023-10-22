package types

import (
	"bective/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

func (t *Todo) GetTodos(c *gin.Context) {
	var todos Todos

	fileBytes := utils.JsonToBytes("data/data.json")
	c.JSON(http.StatusOK, todos.FileToTodos(fileBytes))
}

func (t *Todo) CreateTodo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"todos": "You just created a todo",
	})
}

func (t *Todos) FileToTodos(fileByteValue []byte) Todos {
	var todos Todos
	utils.CheckError(json.Unmarshal(fileByteValue, &todos))
	return todos
}

func (t *Todos) TodosToFile() Todos {
	var todos Todos
	return todos
}
