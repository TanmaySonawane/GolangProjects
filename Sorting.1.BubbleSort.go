// bubble sort
package main

import "fmt"

func main() {
	arr := []int{1, 6, 2, 9, 8, 4, 7, 7, 3}
	len1 := len(arr)
	for i := 0; i < len1; i++ {
		for j := i + 1; j < len1; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
}
