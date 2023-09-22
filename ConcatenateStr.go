//concatenates two strings in a different way: for example aabbcc + ccddee would turn into acacbdbdcece
package main

import (
	"fmt"
)

func main() {
	str1 := "OPTIM"
	str2 := "prime"
	str3 := ""
	for i := 0; i < len(str1); i++ {
		str3 += string(str1[i])
		str3 += string(str2[i])
	}
	fmt.Println(str3)
}
