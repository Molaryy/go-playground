package main

import (
	"fmt"
	"os"
)

func runContainer() {
	fmt.Println("Sup")
}

func main() {
	var cmds [1] []string = ["tes"]


	if len(os.Args) < 2 {
		panic("Provide at leas 1 argument")
	}
	fmt.Println(len(os.Args))
	switch os.Args[1] {
	case "run":
		runContainer()
	default:
		panic("Provide at leas 1 argument")
	}
}