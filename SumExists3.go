// are there any 3 consecutive numbers in the given array that add up to 'x'?
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 13, 111, 1111, 11111, 12, 13, 11}
	x := 14
	var flag bool
	cnt := 1
out:
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				cnt++
				if i != k && i != j && j != k {
					if arr[i]+arr[j]+arr[k] == x {
						flag = true
						fmt.Printf("%d + %d + %d = 14\n", arr[i], arr[j], arr[k])
						fmt.Println("Yes")
						break out
					}
				}
			}
		}
	}

	if flag == true {

	}
	fmt.Println(cnt)
}
