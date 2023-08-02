package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func startQuiz(data []byte) {
	arrayFile := strings.Split(string(data), "\n")
	nbCorrectAnswers := 0
	nbIncorrectAnswers := 0

	for i := 0; i < len(arrayFile); i++ {
		cmpResult := strings.Split(arrayFile[i], ",")

		if len(cmpResult) <= 1 {
			continue
		}

		userInput := ""
		fmt.Print(cmpResult[0], ": ")
		_, err := fmt.Scan(&userInput)
		check(err)

		if userInput == cmpResult[1] {
			nbCorrectAnswers += 1
		} else {
			nbIncorrectAnswers += 1
		}
	}

	fmt.Println("Correct answers = ", nbCorrectAnswers)
	fmt.Println("Incorrect answers = ", nbIncorrectAnswers)
}

func main() {
	useHelp := flag.Bool("fp", false, "Please select a file")
	flag.Parse()
	data, err := os.ReadFile("problems.csv")
	check(err)

	if *useHelp && len(os.Args) > 2 {
		data, err = os.ReadFile(os.Args[2])
		check(err)
		startQuiz(data)
		return
	}
	startQuiz(data)
}
