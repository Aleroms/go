package main

import (
	"fmt"
)

func main() {
	insertAndRemoveFromChannel()

	fmt.Println("----")
	useSelectStatement()

	fmt.Println("----")
	
	commaOkIdiom()
	fmt.Println("----")

	c := make(chan int)
	for i :=0; i<10; i++ {
		go func(){
			for j:=0; j<10;j++{
				c <- j
			}
		}()
		close(c)
	}
	for v:= range c{
		fmt.Println(v)
	}


}
func commaOkIdiom(){
	c := make(chan int)

	go func(){
		c <- 1
	}()
	v, ok := <- c 
	fmt.Println(v, ok)

	close(c)

	v, ok = <- c
	fmt.Println(v,ok)
}
func useSelectStatement(){
	q := make(chan int)
	c := func(q chan<- int)<-chan int{
		c := make(chan int)

		go func(){
			for i := 0; i < 10; i++ {
				c <- i
			}
			q <- 1
			close(c)
		}()

		return c
	}(q)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
func insertAndRemoveFromChannel(){
	// creates a channel
	c := make(chan int)

		
	go func(){
		// inserts onto channel
		c <- 42
		c <- 43

	}()

	// takes off channel
	fmt.Println(<-c)
	fmt.Println(<-c)
}