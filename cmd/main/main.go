package main

import (
	"fmt"
	"log"

	Lab1 "github.com/witchcraft1/architecture-course-laba1"
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
