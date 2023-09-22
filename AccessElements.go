// manually traversing array to access a particular element
package main

import (
	"fmt"
)

func TraverseArray(arr1 []int, i int) {
	for k := 0; k <= i; k++ {
		if k == i {
			fmt.Println(arr1[k])
			break
		}
	}
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	i := 19
	TraverseArray(arr1, i)
}
