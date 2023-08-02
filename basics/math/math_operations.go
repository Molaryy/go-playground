package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

type Func func()

func randomNumber(number int) {
	fmt.Println(rand.Intn(number))
}

func squareRootNumber(number int) {
	fmt.Println(rand.Intn(number))
}

func justPi() {
	fmt.Println(math.Pi)
}

func main() {
	argsLen := len(os.Args)
	var funcSlice []Func

	if argsLen < 2 {
		fmt.Println("Usage:\n\n\tgo run math_operations.go [mathOperators] [argument] ...")
		fmt.Println("\n\tmathOperators: \n \t    --------------------------")
		fmt.Println("\t(1) | random number - 1 args |")
		fmt.Println("\t(2) | square root   - 1 args |")
		fmt.Println("\t(3) | pi number     - 0 args |")
		fmt.Println("\t    |----------------------- |")
		return
	}

	if argsLen >= 2 {
		funcSlice = append(funcSlice, func() { justPi() })
	}
	if argsLen == 3 {
		argument := os.Args[2]
		nb, err := strconv.Atoi(argument)

		if err != nil {
			return
		}

		funcSlice = append(funcSlice, func() { randomNumber(nb) })
		funcSlice = append(funcSlice, func() { squareRootNumber(nb) })
	}
	for i, fn := range funcSlice {
		fmt.Println("I = ", i)
		// if i == nb {
		fn()
		// }
	}
}
