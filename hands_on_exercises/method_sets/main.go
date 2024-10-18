package main

import "fmt"

type person struct{}

func (p *person) speak() {
	fmt.Println("I am a human")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
func main() {
	person1 := person{}
	saySomething(&person1)
	person1.speak()
}