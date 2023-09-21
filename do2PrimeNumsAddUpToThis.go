package main

// are there 2 prime numbers that add up to 'a'?
import "fmt"

func main() {
	var a int = 8
	arr := []int{}
	for i := 2; i <= a; i++ {
		isDivisible := false
		for j := 2; j <= (a / 2); j++ {
			if i%j == 0 && i != j {
				isDivisible = true
			}
		}
		if !isDivisible {
			arr = append(arr, i)
		}
	}
	for i := 0; i < len(arr); i++ {
		found := false
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == a {
				fmt.Println(arr[i], arr[j])
				found = true
			}
		}
		if found {
			break
		}
	}
}
