//using a for loop, this code prints all capital alphabet from A to Z
package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Print(string(rune(i)), " ")
	}
}
