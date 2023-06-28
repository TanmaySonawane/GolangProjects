//package main

import (
	"fmt"
	"math"
)

func main() {
	a := 1634
	b := a
	digits := 0
	sum := 0
	for b > 0 {
		digits++
		b /= 10
	}
	b = a
	for b > 0 {
		sum += int(math.Pow(float64(b%10), float64(digits)))
		b /= 10
	}

	if sum == a {
		fmt.Println("It is an Armstrong Number!")
	} else {
		fmt.Println("It is NOT an Armstrong Number!")
	}
}
