// multiply 2 dimensional matricies
package main

import (
	"fmt"
)

func MultiplyMultiDimMatrix(arr1 [][]int, arr2 [][]int) [][]int {
	rows1 := len(arr1)
	cols1 := len(arr1[0])
	rows2 := len(arr2)
	cols2 := len(arr2[0])

	if cols1 != rows2 {
		fmt.Println("Cannot multiply the matrices. The number of columns in arr1 must be equal to the number of rows in arr2.")
		return nil
	}

	arr3 := make([][]int, rows1)
	for i := 0; i < rows1; i++ {
		arr3[i] = make([]int, cols2)
		for j := 0; j < cols2; j++ {
			arr3[i][j] = 0
			for k := 0; k < cols1; k++ {
				arr3[i][j] += arr1[i][k] * arr2[k][j]
			}
		}
	}
	fmt.Println(arr3)
	return arr3
}

func main() {
	arr1 := [][]int{{1, 2, 3}, {4, 5, 6}}
	arr2 := [][]int{{7, 8}, {9, 10}, {11, 12}}
	result := MultiplyMultiDimMatrix(arr1, arr2)
	if result != nil {
		fmt.Println("Multiplication result:")
		for _, row := range result {
			fmt.Println(row)
		}
	}
}
