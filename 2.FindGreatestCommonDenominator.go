// code for finding the greatest common denominator
package main

import "fmt"

func main() {
	var a int = 30
	var b int = 24
	var i int
	for i = a; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			break
		}
	}
	fmt.Println(i)
}
