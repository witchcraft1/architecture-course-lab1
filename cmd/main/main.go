package main

import (
	"fmt"

	Lab1 "github.com/witchcraft1/architecture-course-laba1"
)

func main() {
	// TODO: Get input from the command line, handle errors.
	res, _ := Lab1.PrefixToInfix("^7*23")
	fmt.Println(res)
}