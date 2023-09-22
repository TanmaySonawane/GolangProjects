// code to add complex numbers
package main

import "fmt"

type ComplexNumber struct {
	real, imag float64
}

func AddComplex(c1 ComplexNumber, c2 ComplexNumber) ComplexNumber {
	sum := ComplexNumber{}
	sum.real = c1.real + c2.real
	sum.imag = c1.imag + c2.imag
	return sum
}

func main() {
	var c1 ComplexNumber
	var c2 ComplexNumber

	c1.real = 2
	c1.imag = 3
	c2.real = 4
	c2.imag = 5

	sum := AddComplex(c1, c2)
	fmt.Println("The addition resulted in:", sum.real, "+", sum.imag, "i")
}
