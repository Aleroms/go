# Errors

Official documentation can be found [here](https://pkg.go.dev/errors).

### Error Handling in Go

Error handling in Go is an explicit, structured way to manage failures and unexpected behavior. Instead of exceptions or try/catch mechanisms like in some other languages, Go uses **error values** returned from functions to signal that something went wrong.

---

### Key Points:

- **Errors are values**: In Go, errors are just values of the type `error`, which can be returned by any function.
- A typical function returns two values: the result (if successful) and an error value. If no error occurred, the error value will be `nil`.
- Errors are checked immediately after the function call to decide how to proceed.

---

### Defining and Returning Errors

The `error` type is a built-in interface in Go, and errors are typically returned as the last value from functions. If a function encounters an issue, it returns an error describing the problem.

#### Example:

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
```

### Explanation:

- **`errors.New()`** creates a simple error with a message.
- The `divide` function checks for an error condition (division by zero) and returns an error if the condition is met.
- **Error checking** is done right after calling `divide` to handle the issue appropriately.

---

### Package Errors

Here is the source code for errors.

```go
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

### Handling Errors with Panic and Recover

While Go encourages returning errors as values, sometimes a **`panic`** is used for unrecoverable conditions. **`panic`** stops the normal flow of execution, and **`recover`** can be used to regain control.

#### Example (Panic and Recover):

```go
package main

import "fmt"

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("something went wrong!")
}

func main() {
	riskyFunction()
	fmt.Println("Program continues after panic.")
}
```

### Explanation:

- **`panic`** is called in case of a serious error.
- **`recover`** is used to regain control after a panic, preventing the program from crashing.

---

### Best Practices for Error Handling

1. **Always check for errors**: After calling a function that returns an error, always handle the error immediately.

   ```go
   result, err := doSomething()
   if err != nil {
       // Handle the error
   }
   ```

2. **Return errors up the call stack**: If you cannot handle an error at your current level, return it to the caller.

3. **Use custom errors for more context**: Custom error types help provide more information and make debugging easier.

4. **Avoid panics for regular error handling**: Use `panic` only for truly unrecoverable situations. For normal error handling, return error values instead.

5. **Wrap errors to add context**: Use `fmt.Errorf` to provide context when returning errors from lower-level functions.
