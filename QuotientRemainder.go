// divide numbers and find their quotient and remainder without using the division operator
package main

import "fmt"

func main() {
	var a int = 12
	var i int = 0
	var b int = 6
	for a >= 0 {
		a -= b
		i++
	}
	if a < 0 {
		i--
	}
	fmt.Println("Quotient of a/b =", i)
	fmt.Println("Remainder of a/b =", a+b)
}
