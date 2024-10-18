package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func race(){
	// this following code causes a race condition
	iterator := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			race := iterator
			runtime.Gosched()
			race++
			iterator = race
			fmt.Println(iterator)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:",iterator)
}
func noRaceMutex(){
	// this following code fixes a race condition
	
	var mu sync.Mutex

	iterator := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			race := iterator
			runtime.Gosched()
			race++
			iterator = race
			fmt.Println(iterator)
			wg.Done()
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("end value:",iterator)
}

func noRaceAtomic(){


	var iterator int64
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&iterator,1)
			fmt.Println(atomic.LoadInt64(&iterator))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:",iterator)
}
func main() {
	// race condition
	// race()
	fmt.Println("-------")
	// uses mutex to end race
	// noRaceMutex()
	fmt.Println("-------")
	// noRaceAtomic()

	// print out the operating system and arch
	fmt.Printf("operating system:%v\tarch:%v\n",runtime.GOOS,runtime.GOARCH)

}