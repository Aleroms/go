# WaitGroup

`package sync`
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls `WaitGroup.Add` to set the number of goroutines to wait for. Then each of the goroutines runs and calls `WaitGroup.Done` when finished. At the same time, `WaitGroup.Wait` can be used to block until all goroutines have finished.

To write concurrent code, we add `go` _in front of a function or method call_

```golang
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
```
