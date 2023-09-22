// prints multiplication tables from 1 to the given number 'a'
package main

import "fmt"

func main() {
	var a int = 12
	for i := 1; i <= a; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println()
	}
}
