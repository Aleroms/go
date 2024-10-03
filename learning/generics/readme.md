# Generics

Generics allow you to generalize the tpye of a value in the function signature for the receiving parameters.

```golang
func addI(a,b int) int {
    return a + b
}
func addF(a,b float64) float64 {
    return a + b
}
// these methods can be combined using generics like...
func addT[T int | float64](a,b T) {
    return a + b
}
```

The above examples are for methods. For types it is like...

```golang
type myNumbers interface {
    int | float64
}

func addT[T myNumbers](a,b T){
    return a + b
}
```

type constraints - the method `addT()` is constraining what types this generic function receives.
type set - the interface takes those types.

## Type Alias

You can create an alias to a type as such.

```golang
type myAlias int
var n myAlias = 42
// will not work
fmt.Println(addT(n, 2))
```

This will not work because `myAlias` is not of type `int` even though its underlining type is. In order for `addT()` to accept `myAlias`, the type set must be changes like...

```golang
// specifies that underlining types of int and float64 are also acceptable
type myNumbers interface {
    ~int | ~float64
}
fmt.Println(addT(,n2))
// [output] 44
```

- (https://go.dev/doc/tutorial/generics)[Go Official Tutorial]

## Concrete types vs interface types

**concrete types** is a _type that you can directly instantiate or create a value from_. The type can directly represent a set of values.

```golang
type Employee struct {
    name string
    age int
}
emp := Employee{
    name: "John Doe",
    age: 25
}
```

- **example** - `int, bool string, float64...`
  **interface type** defines a contract but does not represent a specific data layout/instance. It is abstract that represents behavior but now specific values.

```golang
type Reader interface {
    Read(p []byte){n int, err error}
}
```
