// reverses a number and prints it
package main

import "fmt"

func main() {
	var a int = 67890
	for a > 0 {
		fmt.Print(a % 10)
		a /= 10
	}
}
