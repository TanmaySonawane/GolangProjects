// prints a pyramid of stars 
package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = a
	for i := 1; i <= a; i++ {
		for k := (b); k >= 0; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(" * ")
		}
		fmt.Println("")
		b--
	}

	for i := a; i >= 1; i-- {
		for k := 0; k <= b; k++ {
			fmt.Print(" ")
		}
		for j := i; j >= 1; j-- {
			fmt.Print(" * ")
		}
		fmt.Println("")
		b++
	}

}
