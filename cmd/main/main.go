package main

import (
	"fmt"
	Lab1 "github.com/architecture-course-laba1"
	"log"
)

func main() {
	//args := os.Args[1:]
	//res, err := Lab1.PrefixToInfix(args[0])
	res, err := Lab1.PrefixToInfix("+ * 2 2 * - 10 2 5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	fmt.Println(buildVersion)
}
