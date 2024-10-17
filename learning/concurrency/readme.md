# Concurrency

**Concurrency** is about dealing with multiple tasks at the same time, but not necessarily running them simultaneously. It involves structuring a program so that different parts can make progress independently, often through switching between tasks quickly.

#### Key Points:

- **Concurrency** means managing multiple tasks or processes.
- Tasks are interleaved (not necessarily executed at the same time).
- Often involves **goroutines** in Go, which are lightweight threads of execution.
- Focuses on the **structure** of the program, allowing tasks to be started, paused, and resumed without waiting for others to complete.
- Helps in handling I/O-bound operations efficiently (e.g., waiting for user input, network requests).

#### Example (Concurrency in Go):

```golang
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
    // LAUNCHES GOROUTINE
	go task("Task 1")
	go task("Task 2")

	// Give goroutines time to finish
	time.Sleep(time.Second * 3)
}
```

Goroutines run in the same address space, so access to shared memory must be synchronized.

### Parallelism

**Parallelism** is about executing multiple tasks at the exact same time. This is possible when you have multiple processors or CPU cores. In a parallel system, tasks truly run simultaneously, typically on different CPU cores.

#### Key Points:

- **Parallelism** means running multiple tasks at the same time.
- Requires hardware support (like multi-core CPUs).
- Often used for CPU-bound tasks (e.g., complex computations).
- In Go, this can be achieved using goroutines in combination with multiple CPU cores via `runtime.GOMAXPROCS()`.

#### Example (Parallelism in Go):

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func compute(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Printf("Computing on task %d\n", id)
}

func main() {
	runtime.GOMAXPROCS(2) // Set number of threads to 2 for parallelism

	var wg sync.WaitGroup
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go compute(&wg, i)
	}
	wg.Wait()
}
```

### Contrast: Concurrency vs. Parallelism

| **Aspect**               | **Concurrency**                                     | **Parallelism**                                           |
| ------------------------ | --------------------------------------------------- | --------------------------------------------------------- |
| **Definition**           | Managing multiple tasks by interleaving them.       | Running multiple tasks simultaneously.                    |
| **Execution**            | Tasks are not necessarily running at the same time. | Tasks run at the same time (simultaneous execution).      |
| **Focus**                | Structuring a program to handle multiple tasks.     | Maximizing hardware to execute multiple tasks.            |
| **Best for**             | I/O-bound tasks (waiting for input/output).         | CPU-bound tasks (intensive calculations).                 |
| **Hardware Requirement** | Doesn't require multiple processors or cores.       | Requires multiple processors/cores for true parallelism.  |
| **In Go**                | Implemented using goroutines for task management.   | Can be achieved using goroutines with multiple CPU cores. |

### Summary

- **Concurrency** is about structuring tasks to run independently, often by interleaving their execution.
- **Parallelism** is about running tasks simultaneously, typically on multiple CPU cores.
- Concurrency is a way to make programs handle multiple tasks more efficiently, while parallelism fully utilizes hardware capabilities for performance gains.

---

These notes should help clarify the differences between **concurrency** and **parallelism** and their use cases in programming!
