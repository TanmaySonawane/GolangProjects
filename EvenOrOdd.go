// code to find whether a given number is even or odd
package main

import (
	"fmt"
)

func main() {
	var x int = 29483
	if x%2 == 0 {
		fmt.Println("X is an even number!")
	} else {
		fmt.Println("X is an odd number!")
	}
}
