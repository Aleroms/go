package main

import (
	"fmt"
	"math"
)

func foo58() int {
	return 42
}
func foo59(nums ...int) int{
	var agg int
	for _, v := range nums{
		agg += v
	}
	return agg
}
func foo60(){
	fmt.Println("foo60")
}
func bar58() (int,string){
	return 27, "Santiago"
}
func bar59(nums []int) int {
	var agg int
	for _, v := range nums{
		agg += v
	}
	return agg
}

// ex 60
func showDefer(){
	fmt.Println("This method shows how 'defer' defers...")
	fmt.Println("Calling in this order:\n defer foo60()\n bar58() - returns my name and age")
	defer foo60()
	age,name := bar58()
	fmt.Printf("name:%s\tage:%d\n",name,age)
}

//ex 61
type person struct {
	first string
	age int
}
func(p person) speak(){
	fmt.Printf("Hello, I am %s, %d years old\n",p.first,p.age)
}

//ex 62
type square struct{
	length float64
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}
type shape interface {
	area() float64
}
func info62(s shape) string{
	return fmt.Sprintf("The area is %v",s.area())
}

// ex 63
func Add63(a int, b int) int {
	return a + b
}

func main() {
	/*
	create a func foo that returns an int
	create a func foo that returns an int and string
	print out their values
	*/
	ex58a := foo58()
	fmt.Println("ex58a","=",ex58a)
	ex58b_age, ex58b_name := bar58()
	fmt.Printf("My name is %s and I am %d years old.\n",ex58b_name,ex58b_age)

	/*
	create a func foo that takes a variadic parameter of int
	pass a slice of int as a parameter
	return the sum of all values 
	-
	create a func bar that takes in a parameter of []int
	return the sum of all values passed in
	*/
	ex59a := foo59(1,2,3)
	fmt.Println("ex59a","=",ex59a)
	x59b := []int{4,5,6}
	ex59b := bar59(x59b)
	fmt.Println("ex59b","=",ex59b)

	// display defer keyword
	showDefer()

	/*
	create a struct person and attach a method speak
	create a value of that type and call the method
	*/
	ex61 := person{
		first: "Santiago",
		age: 27,
	}
	ex61.speak()

	/*
	create a type square (len float64, widt float64), type circle (radius float64)
	attach a calculate area method to each
	create a type shape that defines an interface that has area()
	create func info which rcv a shape and prints the area
	create a value of square and circle, use func info to print the square and circle
	*/
	ex62a := circle{
		radius: 3.533,
	}
	ex62b := square{
		length: 420,
		width: 69,
	}
	fmt.Printf("circle: %s\n",info62(ex62a))
	fmt.Printf("square: %s\n",info62(ex62b))
	
	// testing in go
	fmt.Println(Add63(3, 4))
	// ex 64
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)
	x := doMath(42, 16, add)
	fmt.Println(x)
	y := doMath(42, 16, subtract)
	fmt.Println(y)

	
	//create and use an anonymous func
	func(){
		fmt.Println("This is an anonymous function called directly")
	}()

	// assign a func to a variable then call it
	ex69 := func(i int){
		fmt.Println("Anonymous func assigned to a variable",i)
	}
	fmt.Printf("variable %T\n",ex69)
	ex69(1)

	/* 
	create a func that returns a func which returns 42
	assign the func to a var and call the func
	*/
	ex70 := func() func()int {
		return func() int {
			return 42
		}
	}()
	fmt.Printf("a func within a func that returns a func int value: %d\n", ex70())

	/*
	pass a func into a func as an argument
	*/
	fmt.Println(printSquare(square71,5))

	// create a func which encloses the scope o a variable
	f := incrementor()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	
}
func incrementor() func()int {
	incrementor := 0
	return func() int {
		incrementor++
		return incrementor
	}

}

func printSquare(f func(int)int, i int) string{
	return fmt.Sprintf("printing the %d^2 from a callback: %d",i,f(i))
}
func square71(n int) int {
	return n * n
}
//calculates two integers given a function that does arithmetic
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}