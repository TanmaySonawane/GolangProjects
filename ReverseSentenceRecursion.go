// reversing a sentence using recursion
package main

import "fmt"

func RevSenRec(a string, i int, b string) string {
	if len(a) == len(b) {
		return b
	}
	b += string(a[i])
	i--
	return RevSenRec(a, i, b)
}

func main() {
	a := "Hello World!"
	i := len(a) - 1
	b := ""
	fmt.Println(RevSenRec(a, i, b))
}
