// finds the maximum number from the array
package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 4, 4, 77, 35, -33, 00, 5}
	var max int = -1000
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	fmt.Println("Max Number =", max)
}
