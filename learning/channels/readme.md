# Channels

## Channels in GO

**Channels** in Go are a powerful way to communicate between goroutines. They provide a way to send and receive data, allowing goroutines to synchronize and share information without using explicit locks or shared memory. Channels make it easy to coordinate concurrent tasks.

### Key Points:

- Channels are a **typed conduit** through which you can send and receive values of a specific type.
- They can be used to **synchronize** goroutines by blocking them until data is sent or received.
- The channel operations (`send`, `receive`, `close`) can block a goroutine, making them useful for **coordination**.
- sending data to a channel is a blocking operation

---

### Declaring and Using Channels

#### Declaring a Channel:

To declare a channel, use the `chan` keyword followed by the type of data the channel will carry:

```golang
var ch chan int
```

#### Creating a Channel:

Channels are created using the `make` function:

```go
ch := make(chan int)
```

#### Sending and Receiving Data:

- **Send data** into a channel using the `<-` operator:
  ```go
  ch <- 42 // Send the value 42 to the channel
  ```
- **Receive data** from a channel using the `<-` operator:
  ```go
  x := <-ch // Receive data from the channel and store it in x
  ```

---

### Example (Basic Channel Usage):

Here’s an example of a channel used to communicate between two goroutines:

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// Goroutine to send data
	go func() {
		ch <- 42 // Send value to the channel
	}()

	// Receive data from the channel
	value := <-ch
	fmt.Println(value) // Output: 42
}
```

In this example, the `main` goroutine receives a value from the channel that was sent by the anonymous goroutine.

---

### Buffered vs Unbuffered Channels

#### Unbuffered Channels:

An **unbuffered channel** blocks the sender until the receiver is ready, and vice versa.

```go
ch := make(chan int) // Unbuffered channel
```

- Both **send** and **receive** operations will block until the other side is ready.
- This is useful for synchronizing goroutines.

#### Buffered Channels:

A **buffered channel** allows sending and receiving to happen independently up to a specified buffer size.

```go
ch := make(chan int, 3) // Buffered channel with capacity 3
```

- You can send up to 3 values before the channel blocks.
- Receiving from the channel is also non-blocking as long as there are values in the buffer.

#### Example (Buffered Channel):

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2) // Buffered channel with a capacity of 2

	ch <- 1
	ch <- 2
	fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
}
```

In this example, you can send two values before the channel blocks, and then you can receive them later.

### Directional Channels

**Directional channels** in Go restrict the direction in which data can be sent or received. They allow you to specify whether a channel can only be used to **send** or **receive** values in specific parts of your code, improving code safety and clarity.

---

### Types of Directional Channels

1. **Send-Only Channel (`chan<-`)**:

   - Can only be used to **send** data.
   - Useful for restricting a goroutine to only sending values into a channel.

   ```go
   func sendOnly(ch chan<- int) {
       ch <- 42 // Can only send data
   }
   ```

2. **Receive-Only Channel (`<-chan`)**:

   - Can only be used to **receive** data.
   - Useful when a goroutine should only receive values from a channel.

   ```go
   func receiveOnly(ch <-chan int) {
       value := <-ch // Can only receive data
       fmt.Println(value)
   }
   ```

---

### Example (Using Directional Channels):

```go
package main

import "fmt"

func sendOnly(ch chan<- int) {
    ch <- 42 // Send data
}

func receiveOnly(ch <-chan int) {
    fmt.Println(<-ch) // Receive data
}

func main() {
    ch := make(chan int)

    go sendOnly(ch)
    receiveOnly(ch) // Output: 42
}
```

### Range

A range loops over a channel until it is done. It stops reading from a channel when the channel is closed.

```golang
func main(){
    c := make(chan int)

    // send
    go foo(c)

    // receive
    for v := range c {
        fmt.Println(v)
    }

}

// sender
func foo(c chan <- int){
    for i := 0; i < 100; i++ {
        c <- 42
    }
    close(c)
}

```

The `close()` closes the channel. If a channel is not closed while a receiver is reading from it, the receiver may block indefinitely, waiting for more data that will never arrive.

#### deadlock

**deadlock**, where a goroutine is stuck waiting for data that will never come.

```
fatal error: all goroutines are asleep - deadlock!
```

Here are notes on using the **`select` statement** for pulling values off of channels in Go:

---

### Select Statement in Go

The **`select` statement** in Go is used to wait on multiple channel operations (send/receive) and proceed with whichever one is ready. It allows a goroutine to react to data from multiple channels in a non-blocking way.

---

### Key Points:

- **`select`** listens on multiple channels and proceeds with the first one that is ready.
- It's similar to a `switch` statement but specifically designed for handling channel operations.
- Useful for implementing **time-outs**, **waiting for multiple channels**, or ensuring **non-blocking** behavior.
- If none of the channels are ready, the `select` blocks until one of them can proceed.

---

### Example (Pulling Values off a Channel with `select`):

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutines to send data after a delay
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "data from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "data from ch2"
	}()

	// Use select to pull values from whichever channel is ready
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1) // Will execute if ch1 is ready
	case msg2 := <-ch2:
		fmt.Println(msg2) // Will execute if ch2 is ready
	}
}
```

### Explanation:

- **`select`** blocks until one of the channels (`ch1` or `ch2`) is ready to send a value.
- The first channel to send data will trigger its corresponding case, allowing the program to proceed with that value.
- In this example, since `ch2` sends data first (after 1 second), it prints `"data from ch2"`.

---

### Select with a Default Case:

You can add a **`default` case** to make `select` non-blocking. This case executes if no channel is ready.

#### Example (Non-blocking `select`):

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No data received, moving on")
}
```

- **`default`** ensures that the program doesn’t block if no channels are ready.

---

### Select with Timeouts:

You can use `select` with the **`time.After()`** function to implement timeouts while waiting for channel data.

#### Example (Select with Timeout):

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
case <-time.After(2 * time.Second):
    fmt.Println("Timeout!")
}
```

- This example waits for a message from `ch` for up to 2 seconds. If no data arrives, the timeout case executes.

#### Comma OK idiom

It is also possible to use the comma ok idiom to check whether a channel has been closed.

```golang
func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c

	fmt.Println(v, ok)

	v, ok = <-c

	fmt.Println(v, ok)
}
```

#### Fan in pattern

Fan in takes values from many channels and puts those onto one channel

```go
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/
```

#### Fan out pattern

puts chunks of work onto many goroutines

```go
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
```

#### Context

In Go servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. The set of goroutines working on a request typically needs access to request-specific values such as the identity of the end user, authorization tokens, and the request’s deadline. When a request is canceled or times out, all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using. (AVOID LEAKS)
