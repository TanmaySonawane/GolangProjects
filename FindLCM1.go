// code for finding least common multiple
package main

import "fmt"

func main() {
	var a int = 28
	var b int = 12
	var a1 int = a
	var b1 int = b
	var j int = 0
	var i int = 0
	arr1 := make([]int, 0, 50)
	arr2 := make([]int, 0, 50)

	for i <= a1*b1 {
		arr1 = append(arr1, a)
		a += a1
		arr2 = append(arr2, b)
		b += b1
		i += b1
	}

	for j = 0; j < len(arr1); j++ {
		found := false
		for k := 0; k < len(arr2); k++ {
			if arr1[j] == arr2[k] {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Println(arr1[j])
}
