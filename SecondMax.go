// find the second maximum number in an array
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for m := 0; m <= 8; m++ {
		arr := []int{931, 931, 400, 34}
		max := arr[1]
		secMax := arr[0]
		temp := 0
		if secMax > max {
			temp = max
			max = secMax
			secMax = temp
		}
		for i := 0; i < len(arr); i++ {
			if max == secMax {
				secMax = arr[rand.Intn(len(arr))]
			}
			if arr[i] > max {
				secMax = max
				max = arr[i]
			}
			if arr[i] > secMax && arr[i] < max {
				secMax = arr[i]
			}
		}
		fmt.Println("The Second Max number is: ", secMax)
	}
}
