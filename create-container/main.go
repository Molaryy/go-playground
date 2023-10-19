package main

import (
	"fmt"
	"os"
)

func runContainer() {
	fmt.Println("Sup")
}

func main() {
	var cmds [1]string

	cmds[0] = "run"

	if len(os.Args) < 2 {
		panic("Provide at leas 1 argument")
	}
	fmt.Println(len(os.Args))
	switch os.Args[1] {
	case cmds[0]:
		runContainer()
	default:
		panic("Provide at leas 1 argument")
	}
}
