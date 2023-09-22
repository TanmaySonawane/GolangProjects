// finding the exponent of a number recursively
package main

import "fmt"

func PowRec(a int, b int, d int, c int) int {
	if a == b {
		return c
	}
	a *= d
	c++
	return PowRec(a, b, d, c)
}

func main() {
	var a int = 2
	var d int = 2
	var b int = 64
	var c int = 1
	fmt.Println(PowRec(a, b, d, c))
}
