package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo-go/lib"
	"todo-go/routes/home"
	"todo-go/routes/todos"
)

func startApp(envPath string) {
	port := fmt.Sprint("localhost:", lib.GetEnvFileValue(envPath, "PORT"))

	router := gin.Default()

	// home
	home.GetHomeRoute(router)

	// todos
	todos.TodoHandler(router)

	if err := router.Run(port); err != nil {
		return
	}
}

func main() {
	startApp(".env")
}
