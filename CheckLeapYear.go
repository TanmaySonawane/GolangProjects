// checks if the given year is a leap year
package main

import (
	"fmt"
)

func main() {
	var a int = 2002
	if a%4 == 0 {
		fmt.Println("The year", a, "is a leap year!")
	} else {
		fmt.Println("The year", a, "is NOT a leap year!")
	}
}
