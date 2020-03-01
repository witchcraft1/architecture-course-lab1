package main

import (
	"fmt"
	"log"
	"os"

	Lab1 "github.com/witchcraft1/architecture-course-laba1"
)

func main() {
	expr := "+ * 2 2 - 10 2"
	if len(os.Args) > 1 {
		expr = os.Args[1]
	}
	res, err := Lab1.PrefixToInfix(expr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(buildVersion)
}
