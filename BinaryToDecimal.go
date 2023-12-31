// changes a number from binary to decimal
package main

import (
	"fmt"
	"math"
	"strconv"
)

func BinToDec(a int, sum float64, i int) float64 {
	for a > 0 {
		sum += float64(a%10) * math.Pow(float64(2), float64(i))
		a /= 10
		i++
	}
	fmt.Println(sum)
	return sum
}

func DecToBinary(a int, c string) string {
	for a > 0 {
		c += strconv.Itoa((a % 2))
		a /= 2
	}
	c1 := []rune(c)
	length := len(c1)
	for i := 0; i < length/2; i++ {
		c1[i], c1[length-i-1] = c1[length-i-1], c1[i]
	}
	fmt.Print(string(c1))
	return string(c1)
}

func main() {
	var a int
	var b int
	var sum float64
	var i int
	var c string = ""
	fmt.Println("To Convert Binary to Decimal, select 1. To Convert Decimal to Binary, select 2")
	fmt.Scan(&b)
	fmt.Println("What is your number?")
	fmt.Scan(&a)
	if b == 1 {
		BinToDec(a, sum, i)
	}
	if b == 2 {
		DecToBinary(a, c)
	}
}
