// given a range of numbers from a to be, this code prints out all the armstring numbers that fall within that range
// An Armstring number is one where each digit raised to the power of the number of digits, when added with the rest of the digits 
//(that were also raised to the same power), equals the same number, for example 153 = 1^3+5^3+3^3
package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 10
	var b int = 99000
	for i := a; i <= b; i++ {
		digits := 0
		sum := 0
		var c int = i
		for c > 0 {
			digits++
			c /= 10
		}
		c = i
		for c > 0 {
			sum += int(math.Pow(float64(c%10), float64(digits)))
			c /= 10
		}
		if i == sum {
			fmt.Println(i)
		}
	}
}
