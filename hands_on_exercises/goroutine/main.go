package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//launch two additional goroutines
	//each goroutine needs to print something
	// use waitgroups to make each goroutine finish before progra exits
	
	wg.Add(2)
	go goroutine1()
	go func(){
		fmt.Println("goroutine 2")
		wg.Done()
	}()
	// need this at end so the program does not end 
	wg.Wait()
}

func goroutine1(){
	defer wg.Done()
	fmt.Println("goroutine 1")
}