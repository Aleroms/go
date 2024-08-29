// create package-level & block-level variables
// use var, short declaration, and const keywords
// use variables inside main
package main

import "fmt"

var z int 
const y float32 = 4.201

func main(){

	a := "a string"
	fmt.Printf("The value of z is %v and of type %T\n", z, z)
	fmt.Printf("The value of y is %v and of type %T\n", y, y)
	fmt.Printf("The value of a is %v and of type %T", a, a)
	
	
}