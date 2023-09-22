// adding two, 2 dimensional arrays
package main

import (
	"fmt"
)

func AddMultiDimMatrixes(arr1 [][]int, arr2 [][]int) [][]int {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			arr2[i][j] += arr1[i][j]
		}
	}
	fmt.Println(arr2)
	return arr2
}

func main() {
	arr1 := [][]int{{1, 2}, {3, 4}}
	arr2 := [][]int{{5, 6}, {7, 8}}
	AddMultiDimMatrixes(arr1, arr2)
}
