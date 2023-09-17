// this code is to find the factorial of a number without using recursion

package main

import (
	"fmt"
)

func factA(a int) {
	for i := a - 1; i > 0; i-- {
		a *= i
	}
	fmt.Println(a)
}

func main() {
	var a int = 20
	factA(a)
}
