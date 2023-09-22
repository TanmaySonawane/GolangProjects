// prints the fibonacci sequence using recursion
package main

import "fmt"

func Fibo(n0, n1, total int) int {
	total--
	if total == 0 {
		return 0
	}

	newNum := n0 + n1
	n0 = n1
	n1 = newNum
	fmt.Println(newNum)
	return Fibo(n0, n1, total)

}

func main() {
	n0 := 0
	n1 := 1
	total := 10

	Fibo(n0, n1, total)
}
