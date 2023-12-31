// code that gives you the number of vowels and consonants for a given string
package main

import (
	"fmt"
)

func identifyVowels(s string, V int) int {
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			V++
		}
	}
	fmt.Println(V)
	return V
}

func identifyConsonants(s string, C int) int {
	for i := 0; i < len(s); i++ {
		if s[i] != 'a' && s[i] != 'e' && s[i] != 'i' && s[i] != 'o' && s[i] != 'u' && (s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z') {
			C++
		}
	}
	fmt.Println(C)
	return C
}

func main() {
	s := "ae  io  qwttt 8424*$@( i)"
	b := 0
	V := 0
	C := 0

	fmt.Println("To find the number of vowels, select 1. To find the number of consonants, select 2 \n Only works for small letters")
	fmt.Scan(&b)
	if b == 1 {
		identifyVowels(s, V)
	}
	if b == 2 {
		identifyConsonants(s, C)
	}
}
