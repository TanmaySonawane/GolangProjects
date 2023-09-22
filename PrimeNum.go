// finds if a given number is a prime number or not
package main

import "fmt"

func isPrime(a int) bool {

	if a <= 1 {
		return false
	}

	for i := 2; i <= int(a/2)+1; i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
func main() {

	a := 7

	if isPrime(a) {
		fmt.Println("is a Prime Number")
	} else {
		fmt.Println("womp womp waaa!")
	}
}
