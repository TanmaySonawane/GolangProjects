//Find greatest common denominator using recursion
package main

import "fmt"

func RecursiveGCD(a int, b int, div int) int {
	fmt.Println("div :", div)
	if div == 1 {
		return 1
	}
	if a%div == 0 && b%div == 0 {
		return div
	}
	div--
	return RecursiveGCD(a, b, div)
}

func main() {
	const a int = 16
	const b int = 24
	fmt.Println("GCD is: ", RecursiveGCD(a, b, 16))
}
