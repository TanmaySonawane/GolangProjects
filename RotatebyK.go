// rorates a list by k elements
package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	k := 2
	for i := 0; i < k; i++ {
		arr1 = append(arr1, arr1[i])
		arr1[i] = 0
	}
	for i := k; i < len(arr1); i++ {
		fmt.Print(arr1[i])
	}
}
