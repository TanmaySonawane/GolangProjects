// reads the contents of a file
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
