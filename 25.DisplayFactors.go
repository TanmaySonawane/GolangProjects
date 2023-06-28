// prints out the factors of the given number num
package main

import "fmt"

func main() {
	var num int = 500
	for i := 2; i <= (num / 2); i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
}
