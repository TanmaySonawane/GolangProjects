// finding standard deviation of an array
package main

import (
	"fmt"
	"math"
)

func StdDev(arr []float64, sum float64, sum2 float64) float64 {
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	sum /= float64(len(arr))
	for i := 0; i < len(arr); i++ {
		sum2 += float64(int(math.Pow((sum - arr[i]), 2)))
	}
	sum2 = math.Sqrt(sum2 / float64(len(arr)-1))
	return sum2
}
func main() {
	arr := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var sum2 float64 = 0
	var sum float64 = 0
	fmt.Println(StdDev(arr, sum, sum2))
}
