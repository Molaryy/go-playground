package main

import (
	"fmt"
	"lib"
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

func main() {

	//router := gin.Default()
	port := fmt.Sprint("localhost:", lib.GetEnvFileValue(".env", "PORT"))
	fmt.Println(port)
	//router.Run("localhost:3000")
}
