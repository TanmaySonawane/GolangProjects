//given a range, starting with 'begin' and ending with 'end', this code prints out all the prime numbers in between the range
package main

import "fmt"

func main() {
	var begin int = 3
	var end int = 150
	for i := begin; i <= end; i++ {
		isDivisible := false
		for j := 2; j <= (int(i/2)); j++ {
			if i%j == 0 {
				isDivisible = true
			}
		}
		if !isDivisible {
			fmt.Println(i)
		}
	}
}
