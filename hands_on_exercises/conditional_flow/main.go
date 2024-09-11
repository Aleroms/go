package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250 + 1)
	fmt.Printf("x is %d \n", x)

	if x < 101 {
		fmt.Printf("between 0 and 100")
	} else if x < 201 {
		fmt.Printf("between 101 and 200")
	} else {
		fmt.Printf("between 201 and 250")
	}
}