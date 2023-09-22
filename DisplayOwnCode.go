// code that displays its own code as the output
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("DisplayOwnCode.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
