## Benchmarking

Go’s testing package also supports **benchmarking** to measure the performance of code. Benchmarks are written in `*_test.go` files and use functions starting with `Benchmark`.

#### Key Points:

- Benchmark functions start with `Benchmark` (e.g., `BenchmarkAdd`).
- They use `b *testing.B` to control the benchmark and report timing.
- The loop `for i := 0; i < b.N; i++` runs the code multiple times to average the time taken.

#### Example (Benchmark Test):

```go
package math

import "testing"

func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}
```

Run the benchmark with:

```bash
go test -bench=.
```

### Explanation:

- **`b.N`**: The testing framework sets `b.N` to control how many times the benchmarked code runs.
- Benchmarks provide an average time per operation, which helps evaluate performance.

---

### Examples in Go

**Examples** in Go serve as runnable documentation, showcasing how to use functions and APIs. Example functions begin with `Example`, and they are placed in `*_test.go` files like tests and benchmarks. Examples can be run and checked to ensure they produce the expected output.

#### Key Points:

- Example functions start with `Example` (e.g., `ExampleAdd`).
- The function body demonstrates how to use code from the package.
- Expected output is included in a comment at the end of the function.

#### Example (Documentation Test):

```go
package math

import "fmt"

func ExampleAdd() {
    result := Add(2, 3)
    fmt.Println(result)
    // Output: 5
}
```

Run examples as part of `go test` to ensure they work as expected and match the output specified.

#### Explanation:

- **`// Output:`** comments specify the expected output.
- Example functions help document code usage and verify that documentation matches actual behavior.

---

### Summary

| **Concept**      | **Purpose**                            | **Function Prefix** | **Command**               | **Example**                                                  |
| ---------------- | -------------------------------------- | ------------------- | ------------------------- | ------------------------------------------------------------ |
| **Testing**      | Verify code correctness                | `Test`              | `go test`                 | `func TestAdd(t *testing.T) { ... }`                         |
| **Benchmarking** | Measure performance                    | `Benchmark`         | `go test -bench=.`        | `func BenchmarkAdd(b *testing.B) { ... }`                    |
| **Examples**     | Provide usage examples (documentation) | `Example`           | `go test` (checks output) | `func ExampleAdd() { fmt.Println(Add(2, 3)); // Output: 5 }` |
