// code to manually compute the length of a string
package main

import "fmt"

func lenString(s string) int {
	len1 := 0
	for range s {
		len1++
	}
	return len1
}

func main() {
	s := "Hello World!"
	fmt.Println(lenString(s))
}
