// sorts the given two arrays and produces a new array
package main

import "fmt"

func main() {
	arr1 := []int{1, 6, 33, 73}
	arr2 := []int{2, 4, 6, 8}
	arr3 := []int{}
	for i := 0; i < len(arr1); i++ {
		arr3 = append(arr3, arr1[i])
		arr3 = append(arr3, arr2[i])
	}
	
	for i := 0; i < len(arr3); i++ {
		for j := i + 1; j < len(arr3); j++ {
			if arr3[j] < arr3[i] {
				tmp := arr3[i]
				arr3[i] = arr3[j]
				arr3[j] = tmp
			}
		}
	}
	fmt.Println(arr3)
}
