// code that keeps only letters in a string and removes the rest
package main

import (
	"fmt"
)

func RemoveChar(s string, t string) string {
	for i := 0; i < len(s); i++ {
		if (string(s[i]) >= string(rune(65)) && string(s[i]) <= string(rune(90))) || (string(s[i]) >= string(rune(97)) && string(s[i]) <= string(rune(122))) {
			t += string(s[i])
		}
	}
	return t
}

func main() {
	s := "abcde1234!@#$%^&*fgh"
	t := ""
	fmt.Println(RemoveChar(s, t))
}
