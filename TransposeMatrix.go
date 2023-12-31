// finding the transpose of a matrix
package main

import (
	"fmt"
)

func TransposeMatrix(arr1 [][]int) [][]int {
	rows := len(arr1)
	cols := len(arr1[0])
	arr2 := make([][]int, rows)
	for i := 0; i < rows; i++ {
		arr2[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			arr2[i][j] = arr1[j][i]
		}
	}
	return arr2
}

func main() {
	arr1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Original Matrix: ")
	for _, row := range arr1 {
		fmt.Println(row)
	}
	fmt.Println("Transposed Matrix:")
	for _, row := range TransposeMatrix(arr1) {
		fmt.Println(row)
	}
}
