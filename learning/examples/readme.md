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

### Summary

| **Concept**      | **Purpose**                            | **Function Prefix** | **Command**               | **Example**                                                  |
| ---------------- | -------------------------------------- | ------------------- | ------------------------- | ------------------------------------------------------------ |
| **Testing**      | Verify code correctness                | `Test`              | `go test`                 | `func TestAdd(t *testing.T) { ... }`                         |
| **Benchmarking** | Measure performance                    | `Benchmark`         | `go test -bench=.`        | `func BenchmarkAdd(b *testing.B) { ... }`                    |
| **Examples**     | Provide usage examples (documentation) | `Example`           | `go test` (checks output) | `func ExampleAdd() { fmt.Println(Add(2, 3)); // Output: 5 }` |
