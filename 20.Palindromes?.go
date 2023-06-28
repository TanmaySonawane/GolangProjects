//checks if the given variable 'num' is a number palindrome, meaning it is the same forwards and backwards
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 12344321
	a := num
	var t string = ""
	for a > 0 {
		t += strconv.Itoa((a % 10))
		a /= 10
	}
	s := strconv.Itoa(num)
	if s == t {
		fmt.Println("Your number is a palindrome")
	} else {
		fmt.Println("Your number is NOT a palindrome")
	}
}
