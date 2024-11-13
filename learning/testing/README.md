# Testing in Go

## Tests

Go’s standard library includes a built-in **testing framework** to create unit tests. Tests are written as regular Go functions and placed in `*_test.go` files within the same package as the code being tested.

[documentation](https://pkg.go.dev/testing)

#### Key Points:

- Test functions start with `Test` (e.g., `TestAddition`).
- Tests should use `t *testing.T` to log errors and manage test failures.
- Test files end with `_test.go` and are run using `go test`.

#### Example (Unit Test):

Given a simple addition function:

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}
```

A test for `Add` would look like:

```go
// math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

Run the test with:

```bash
go test
```

#### Additional Testing Tips:

- Use `t.Error` or `t.Errorf` to log test failures.
- Use `t.Fatal` or `t.Fatalf` if the test should stop immediately upon failure.
