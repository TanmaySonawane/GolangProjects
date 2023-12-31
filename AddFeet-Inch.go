// when given two distances in feet and inches, this code correctly adds them and give the final answer in feet and inches
package main

import "fmt"

type distance struct {
	feet   int
	inches float64
}

func main() {
	var distance1 distance
	var distance2 distance

	fmt.Println("For Distance 1:")
	fmt.Println("Please enter feet:")
	fmt.Scan(&distance1.feet)
	fmt.Println("Please enter inches:")
	fmt.Scan(&distance1.inches)

	fmt.Println("For Distance 2:")
	fmt.Println("Please enter feet:")
	fmt.Scan(&distance2.feet)
	fmt.Println("Please enter inches:")
	fmt.Scan(&distance2.inches)

	distance1.feet += distance2.feet
	distance1.inches += distance2.inches
	for distance1.inches >= 12 {
		distance1.feet++
		distance1.inches -= 12
	}
	fmt.Println("Adding the two distances, Feet:", distance1.feet, "Inches:", distance1.inches)
}
