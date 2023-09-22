// write a sentence to a file using io
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err := os.Create("data.txt")
	val := "I am Batman\n"
	data := []byte(val)
	err3 := ioutil.WriteFile("data.txt", data, 0)
	if err3 != nil {
		log.Fatal(err3)
	}

	_, err2 := f.WriteString("I am Iron Man\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	words := []string{"I", "am", "Spider-", "Man"}

	for _, word := range words {

		_, err := f.WriteString(word + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	const name, age = "Pramey ", 34

	n, err := fmt.Fprintln(f, name, "is", age, "years old.")
	fmt.Println(n, "bytes written")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("Done!")
}
