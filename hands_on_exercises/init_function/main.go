package main

import (
	"fmt"
	"math/rand"
)

func init(){
	fmt.Println("This is where initialization for my program occurs")
	// A niladic function is a special type of function that takes no arguments 
	// and is evaluated immediately when it is encountered in a statement.
}
func main() {
	x := rand.Intn(250 + 1)
	fmt.Printf("x is %d \n", x)

	if x < 101 {
		fmt.Println("between 0 and 100")
	} else if x < 201 {
		fmt.Println("between 101 and 200")
	} else {
		fmt.Println("between 201 and 250")
	}

	switch {
	case x < 101:
		fmt.Println("between 0 and 100")
	case x < 201:
		fmt.Println("between 101 and 200")
	case x < 251:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("out of scope")
	}
}