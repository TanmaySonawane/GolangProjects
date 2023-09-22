// turns any english sentence into pig latin. In Pig latin, the basic rule is to switch the first consonant 
// to the end of the term and then adding suffix “ay” to form a new word
// For instance, the word 'pig' would become igp+ay which becomes igpay
package main

import "fmt"

func main() {

	s := "The quick brown fox jumped over the fence "
	firstChFound := false
	firstChOfWord := ""
	wordRemainChs := ""

	for i := 0; i < len(s); i++ {
		ch := string(s[i])
		if firstChFound == false && ch != " " {
			firstChOfWord = ch
			firstChFound = true
			continue
		}

		if ch == " " {
			firstChFound = false
			pigLatinWord := wordRemainChs + firstChOfWord + "ay "
			fmt.Printf(pigLatinWord)
			wordRemainChs = ""
			firstChOfWord = ""
		}
		wordRemainChs += ch
	}
}
