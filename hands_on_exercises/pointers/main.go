package main

import "fmt"

type young interface{
	isYoung() bool
}
type boy struct{
	age uint8
}
func (b boy) isYoung() bool {
	return true
}
type oldMan struct {
	age uint8
}
func(m oldMan) isYoung() bool {
	return false
}

type person struct {
	first string
}


func main() {
	// create a value and assign it to a variable; print the address
	ex74 := 3
	fmt.Printf("The memory location of ex74 is %v, the value is %v\n",&ex74, ex74)
	// dereference the addres of a variable
	ex75a := 42
	ptr_ex75a := &ex75a
	fmt.Printf("Dereferencing ex75a: %v\nType is %T\n",*ptr_ex75a,ptr_ex75a)
	
	/*
	implement an interface and make a value/pointer semantics
	*/
	ex76a := boy{age: 10}
	ex76b := oldMan{age: 32}
	ex76a.age = 11
	fmt.Println(ex76a.isYoung(),ex76b, ex76a)

	/*
	create two functions to change the field in a struct
	use value semantics (returns a value of some type)
	use pointer semantics (returns nothing)
	*/
	ex77a := person{first: "James"}
	fmt.Printf("%#v\n",ex77a)
	ex77a = changeName(ex77a, "Santiago")
	fmt.Printf("%#v\n",ex77a)

	ex77b := person{first: "James"}
	fmt.Printf("%#v\n",ex77b)
	changeNamePtr(&ex77b, "Santiago")
	fmt.Printf("%#v\n",ex77b)
	
	
	
}
// value semantics
func changeName(p person, s string) person {
	p.first = s
	return p
}
// pointer semantics
func changeNamePtr(p *person, s string) {
	p.first = s
}