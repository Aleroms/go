package main

import (
	"math/rand"
)

func main() {
	x,y := rand.Intn(10), rand.Intn(10)

	 if x < 4 && y < 4 {
		println("x and y are both less than 4")
	} else if x > 6 && y > 6 {
		println("x and y are both greater than 6")
	} else if x >= 4 && x <= 6 {
		println("x is greater than or equal to 4 and less than or equal to 6")
	} else if y != 5 {
		println("y is not 5")
	} else {
		println("no case met")
	}

	switch {
	case x < 4 && y < 4:
		println("x and y are both less than 4")
	case x > 6 && y > 6:
		println("x and y are both greater than 6")
	case x >= 4 && x <= 6:
		println("x is greater than or equal to 4 and less than or equal to 6")
	case y != 5:
		println("y is not 5")
	default:
		println("no case met")
	}
}