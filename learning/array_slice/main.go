package main

import "fmt"

func array() {
	// array literal
	a := [3]int {4,3,2}
	fmt.Println(a)

	b := [...]string {"Hello", "Gophers"}
	fmt.Println(b)

	var c [2]int
	c[0] = 7
	c[1] = 8

	fmt.Println(c)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",c)

	// Hands on exercise 
	/*
	Use an array literal to store these elements in an array and let the compiler determine
	the length of the array, then also print out the array and the length of the array
	*/
	d := [...]string{
		"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", 
		"Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
		 "Browned ButterCookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie",
		  "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)",
		   "Lavender Lemon Honey (GF)", "Lemon Bar", "MadagascarVanilla (GF)", "Matcha (GF)", 
		   "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
		    "Non-Dairy SunbutterCinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", 
			"RaspberrySorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry LemonadeSorbet (GF, V)", 
			"Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)",
	}
	fmt.Println(d)
	fmt.Println(len(d))
	fmt.Printf("%T\n",d)

	// vs code shortcut 
	// alt click - selects multiple by placing cursor at click
	// alt+shift+mouse_drag - places cursor at the front 
	// alt+shift+i+drag - places cursor at the end of each line selected 
	
}

func slice() {
	//slice literal - don't add size basically
	xs := []string{"hello", "world"}
	fmt.Println(xs)

	// Hands on exercise
	// create a slice literal to store elements and print out slice & length 
	// then, use for range to loop over slice 
	exercise41 := []string{
		"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", 
		"Bittersweet Chocolate (GF)",
	}
	fmt.Println(exercise41)
	fmt.Println(len(exercise41))
	fmt.Printf("%T\n",exercise41)
	for _, value := range exercise41 {
		fmt.Printf("%v\t", value)
	}
	fmt.Println("-----------------")
	fmt.Println(exercise41[0])
	fmt.Println(exercise41[1])
	fmt.Println(exercise41[2])
	
	xi := []int{42,43,44}
	fmt.Println(xi)
	xi = append(xi, 45)
	fmt.Println(xi)
	xi = append(xi, 46, 47)
	fmt.Println(xi)
	fmt.Printf("length of xi is %v\n",len(xi))

	//slicing a slice
	//[inclusive:exclusive]
	fmt.Printf("xi - %v\n",xi[0:4])
	
	//[:exclusive]
	fmt.Printf("xi - %v\n",xi[:2])
	
	//[inclusive:]
	fmt.Printf("xi - %v\n",xi[3:])
	
	//[:] - everything
	fmt.Printf("xi - %v\n",xi[:])

	// kind of like the spread operator in JS
	fmt.Println(xi)
	xi = append(xi[:1], xi[2:]...)
	fmt.Println(xi)

	// make function
	si := []string{"A","B","C"}
	fmt.Println(si)

	di := make([]int,0,10)
	fmt.Println(di)
	fmt.Println(len(di))
	fmt.Println(cap(di))
	di = append(di, 0,1,2,3,4,5,6,7,8,9)
	fmt.Println(di)


	// multidimensional array
	a := []int{1,2,3,4,5}
	b := []int{6,7,8,9,0}
	xxab := [][]int{a,b}
	fmt.Println(xxab)
}

func main(){
	array()
	slice()
}

