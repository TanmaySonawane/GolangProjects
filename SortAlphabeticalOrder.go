// sorts the elements in the string alphabetically
package main

import (
	"fmt"
)

func SortAplhabetically(arr []string) []string {
	var temp string
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

func main() {
	arr := []string{"daniel", "max", "harry", "abraham", "zendaya", "nicholas"}
	fmt.Println(SortAplhabetically(arr))
}
