// created a struct to store student information
package main

import "fmt"

type student struct {
	name, city  string
	age, rollNo int
}

func main() {
	std1 := student{
		name:   "Jamal",
		age:    45,
		city:   "dholakpur",
		rollNo: 57,
	}

	var std2 student
	std2 = student{"Ramesh", "Dencity", 72, 360}

	fmt.Println("Student 1", "name:", std1.name, "\n age:", std1.age, "\ncity:", std1.city, "\nroll No.", std1.rollNo)
	fmt.Println("Student 2", std2)

	fmt.Println("Are there more students? Y = yes, N = no")
	N := ""
	name := ""
	age := 0
	city := ""
	rollNo := 0
	fmt.Scan(&N)
	if N != "Y" {
		return
	}

	students := []student{}

	for N == "Y" {
		fmt.Println("Please enter name")
		fmt.Scan(&name)
		fmt.Println("Enter Age")
		fmt.Scan(&age)
		fmt.Println("Enter City Name")
		fmt.Scan(&city)
		fmt.Println("Enter Roll No")
		fmt.Scan(&rollNo)

		newStudent := student{
			name:   name,
			age:    age,
			city:   city,
			rollNo: rollNo,
		}

		students = append(students, newStudent)

		fmt.Println("Are there more students? Y = yes, N = no")
		fmt.Scan(&N)
	}

	for i, s := range students {
		fmt.Printf("Student %d:\n", i+1)
		fmt.Println("Name:", s.name)
		fmt.Println("Age:", s.age)
		fmt.Println("City:", s.city)
		fmt.Println("Roll No.:", s.rollNo)
		fmt.Println()
	}
}
