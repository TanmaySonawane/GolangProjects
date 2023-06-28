package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Println("What is the first number?")
	fmt.Scan(&a)
	var c string
	fmt.Println("What operation do you want to perform?")
	fmt.Scan(&c)
	var b int
	fmt.Println("What is the second number?")
	fmt.Scan(&b)

	switch c {
	case "+":
		fmt.Println(a, "+", b, "=", a+b)
	case "-":
		fmt.Println(a, "-", b, "=", a-b)
	case "*":
		fmt.Println(a, "*", b, "=", a*b)
	case "/":
		fmt.Println(a, "/", b, "=", a/b)
	default:
		fmt.Println("Invalid")
	}
}
