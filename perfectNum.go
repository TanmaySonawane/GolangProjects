// checks if the given number is a perfect number
// a perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself
package main

import "fmt"

func Perfect(a int, myList []int) bool {
	for i := 2; i <= int(a/2)+1; i++ {
		if a%i == 0 {
			myList = append(myList, i)
		}
	}
	return Perfect2(a, myList)
}

func Perfect2(a int, myList []int) bool {
	sum := 0
	for _, num := range myList {
		sum += num
	}
	if sum+1 == a {
		return true
	}
	return false
}

func main() {
	myList := make([]int, 0, 100)
	a := 6
	if Perfect(a, myList) {
		fmt.Println(a, "is a Perfect Number")
	} else {
		fmt.Println(a, "is not a Perfect Number")
	}
}
