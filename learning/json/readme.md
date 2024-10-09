# JSON

## JSON marshal

Marshal in Go is the process of converting your Go data structures (like structs, maps, etc.) into JSON format. This is helpful when you need to send data as JSON in an API request or save it in a file.

```golang
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{Name: "Alice", Age: 30}

	// Convert struct to JSON (marshal)
	jsonData, _ := json.Marshal(person)

	// Print the JSON data as a string
	fmt.Println(string(jsonData))
    // Output: {"name":"Alice","age":30}
}
```

In this example, Go takes the person struct and turns it into JSON. Remember that the properties in the struct must be public, so capitalize them.

You can also encode data directly to a stream (like the console or a file) using `json.NewEncoder`.

```golang
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Encode users to JSON and write directly to os.Stdout (the console)
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}
```

`json.NewEncoder(os.Stdout)` creates a new JSON encoder to write to os.Stdout (the console in this case).
`.Encode(users)` converts the users slice into JSON format and directly writes it to the output stream.
This is useful when you want to immediately write the JSON output to a stream without manually handling the marshaled result.

## JSON Unmarshal

Unmarshal is the opposite of marshal. It converts JSON data back into Go data structures. This is useful when you receive JSON from an API and need to work with it in Go.

```golang
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := []byte(`{"name":"Alice","age":30}`)

	// Create an empty Person struct
	var person Person

	// Convert JSON to struct (unmarshal)
	_ = json.Unmarshal(jsonData, &person)

	// Print the Go struct
	fmt.Println(person)  // Output: {Alice 30}
}
```

In this example, JSON data is converted back into a Go Person struct.

Summary

- Marshal: Go struct ➡️ JSON
- Unmarshal: JSON ➡️ Go struct

These processes are key when working with APIs or files in Go!

## Writer Interface

The `Writer` interface in Go is used for writing data to an output, such as a file, buffer, or network connection. It’s part of the io package and is very simple: if a type has a `Write` method, it implements the Writer interface.

```golang
package main

import (
	"fmt"
	"os"
)

func main() {
	message := []byte("Hello, Go!")

	// os.Stdout implements the Writer interface
	os.Stdout.Write(message)  // Output: Hello, Go!
}
```

In this example, `os.Stdout` is an output stream that implements the `Writer` interface, and `Write` sends the message to the console.

## Sort

Go provides a `sort` package that allows you to sort slices of basic types like int, float64, and string.

```golang
package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 3, 8, 1, 2}

	// Sort the slice of integers
	sort.Ints(numbers)

	fmt.Println(numbers)  // Output: [1 2 3 5 8]
}
```

## Custom Sort

In Go, you can sort complex types by implementing the sort.Interface. This requires three methods: Len(), Swap(), and Less(). In this example, we’ll sort a slice of Person structs by age.

```golang
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

// ByAge implements the sort.Interface for []Person based on the Age field.
type ByAge []Person

// Len returns the length of the slice
func (a ByAge) Len() int { return len(a) }

// Swap swaps the positions of two elements in the slice
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less compares two elements and returns true if the first is less than the second
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("Before sorting:", people)

	// Sort the slice by age
	sort.Sort(ByAge(people))

	fmt.Println("After sorting by age:", people)
}
```

1. Struct Definition: We define a Person struct with First (name) and Age fields.
2. ByAge Type: A new type, ByAge, is created as a slice of Person. This will allow us to implement custom sorting behavior.
3. Len Method: Returns the length of the slice.
4. Swap Method: Swaps two elements in the slice.
5. Less Method: Defines the sorting logic. In this case, it compares the Age field to sort in ascending order.
6. sort.Sort: We pass the ByAge type to sort.Sort() to sort the slice of Person by age.

## bcrypt

bcrypt is a password hashing function in Go’s golang.org/x/crypto/bcrypt package. It's used to securely hash and verify passwords.

```golang
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "mysecretpassword"

	// Hash the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	fmt.Println(string(hashedPassword))  // Output: hashed version of the password

	// Verify the password
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte("mysecretpassword"))
	if err != nil {
		fmt.Println("Password does not match!")
	} else {
		fmt.Println("Password matches!")
	}
}
```

In this example, the password is hashed using bcrypt, and then we check if the plain text password matches the hashed password.
