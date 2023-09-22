// swapping two variables using a third variable temp
package main

import "fmt"

func main() {
   var a int = 3
   var b int = 2

   swap(&a, &b)

   fmt.Printf("a =", a ,"& b =", b)
}
func swap(x *int, y *int) {
   var temp int
   temp = *x    
   *x = *y    
   *y = temp  
}
