// find out how often a particular character appears in a string
package main

import "fmt"

func FreqChar(s string, a string, c int) int {
	for i := 0; i < len(s); i++ {
		if a == string(s[i]) {
			c++
		}
	}
	return c
}

func main() {
	s := "abbcccddddeeeee"
	a := "d"
	c := 0
	fmt.Println(FreqChar(s, a, c))
}
