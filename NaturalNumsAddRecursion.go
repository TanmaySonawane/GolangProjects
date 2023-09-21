// Recursively adds a number to its previous number
package main

import "fmt"

func RecursiveAdd(a int) int {
	if a == 0 {
		return 0
	}
	return a + RecursiveAdd(a-1)
}

func main() {
	var a int = 3
	fmt.Println(RecursiveAdd(a))
}
