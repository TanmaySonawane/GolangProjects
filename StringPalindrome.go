// tells whether a string is a palindrome or not
package main

import (
	"fmt"
)

func main() {
	s := "hannnah"
	i := 0
	j := len(s) - 1
	palindrome := true
	for i != j {
		if s[i] != s[j] {
			palindrome = false
		}
		i++
		j--
		if i+1 == j {
			if s[i] != s[j] {
				palindrome = false
			}
			break
		}
	}
	if palindrome {
		fmt.Println(s, "is a palindrome")
	} else {
		fmt.Println(s, "is NOT a palindrome")
	}
}
