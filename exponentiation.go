// raising a to the power b without using power or times operator
package main

import (
	"fmt"
)

func main() {
	a := 2
	b := 3
	newA := a
	mul := 0
	for i := 0; i < b-1; i++ {
		for j := 0; j < a; j++ {
			mul += newA
		}
		newA = mul
		mul = 0
	}

	fmt.Println(newA)
}
