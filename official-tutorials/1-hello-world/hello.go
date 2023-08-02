package main

import "fmt"

import (
	"rsc.io/quote"
	"rsc.io/sampler"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(sampler.Hello())
}
