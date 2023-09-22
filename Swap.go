// swap 2 variables WITHOUT using a third variable
package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 0
	var c int = 20
	b = a
	a = c
	c = b
	fmt.Println("a = ", a, "c = ", c)
}
