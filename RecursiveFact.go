// use recursion to find factorial
package main

import "fmt"

func RecursiveFactorial(a int) int {
	if a == 0 || a == 1 {
		return a
	}
	return a * RecursiveFactorial(a-1)
}

func main() {
	a := 4
	fmt.Println(RecursiveFactorial(a))
}
