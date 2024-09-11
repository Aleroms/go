package main

import (
	"fmt"
	"math/rand"
)

// THERE IS NO WHILE LOOP
func main() {
	// part one
	for i := 0; i < 100; i++ {
		fmt.Printf("%d\t", i)
	}
	// part two

	for a:= 0; a < 100; a++ {
		x,y := rand.Intn(10), rand.Intn(10)
		switch {
		case x < 4 && y < 4:
			fmt.Printf("x and y are both less than 4 \t")
		case x > 6 && y > 6:
			fmt.Printf("x and y are both greater than 6 \t")
		case x >= 4 && x <= 6:
			fmt.Printf("x is greater than or equal to 4 and less than or equal to 6 \t")
		case y != 5:
			fmt.Printf("y is not 5 \t")
		default:
			fmt.Printf("no case met \t")
		}
	}

	// part three
	x := 100

	for x > 0 {
		fmt.Printf("x is %v\t", x)
		x--
	}

	// part four
	x = 10
	for x > 0 {
		x--
		if x == 5 {
			fmt.Println("break")
			break
		}
	}
	fmt.Println(x)
	
	// part five 
	for z:= 0; z<30; z++{
		if z % 2 == 0 {
			continue
		}
		fmt.Printf("%v ",z)
	}
	fmt.Println()

	// nested for loops
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i=%v, j=%v\n",i,j)
		}
	}

	// for range loop
	xi := []int{42, 43, 44, 45, 46, 47}

	for i,v := range xi {
		fmt.Printf("index=%v; value=%v\t",i,v)
	}
	fmt.Println()
	// for range loop 2
	m := map[string] int {
		"james": 42,
		"moneypenny": 32,
	}
	for key, value := range m {
		fmt.Printf("key=%v value=%v;\t", key, value)
	}
	age := m["james"]
	fmt.Println(age)
	fmt.Println(m["Q"])

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, value of v",v)
	}

	fmt.Println()

	// there is an easier way to do this
	// using statement statement idiom
	// for i:=0; i< 100; i++{
		
	// 	xx := rand.Intn(5)
	// 	if xx == 3 {
	// 		fmt.Println("x is 3")
	// 	}
	// }
	for i:=0; i<100; i++{
		if xx:= rand.Intn(5); xx == 3 {
			fmt.Println("x is 3")
		}
	}
}