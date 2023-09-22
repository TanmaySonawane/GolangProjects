// Asks the user to guess a number and checks if it matches the written number, keeps track of user's tries
package main

import (
	"fmt"
)

func main() {
	var correct int = 12 // make it anything other than 45
	var guess int
	var tries int
	arr := []int{}
	for guess != correct {
		fmt.Println("Please Guess the Number, Enter '45' to exit")
		if guess == 45 {
			break
		}
		fmt.Scan(&guess)
		if guess > correct {
			fmt.Println("Too High")
		}
		if guess < correct {
			fmt.Println("Too Low")
		}
		found := false
		for i := 0; i < len(arr); i++ {
			if arr[i] == guess {
				found = true
				break
			}
		}
		if found {
			fmt.Println("You already guessed that number")
		} else {
			arr = append(arr, guess)
			tries++
		}
	}
	if guess == correct {
		fmt.Println("Hooray! You Won!")
		fmt.Println("Number of tries:", tries)
	}
}
