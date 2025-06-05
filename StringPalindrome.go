// tells whether a string is a palindrome or not
package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	s := "hannnah"

	if isPalindrome(s) {
		fmt.Println(s, "is a palindrome")
	} else {
		fmt.Println(s, "is NOT a palindrome")
	}
}
