# Functions in Go

## Syntax of functions in GO

```golang
func(r receiver) identifier(p parameters)(return(s)){...}
```

**everything in GO is passed by value** when you pass in an argument, it is beign assigned to a variable. There can be multiple returns as well.

```golang
// no parameters, no returns - essentially a void function
func foo(){
    fmt.Println("I am from foo")
}

// one parameter, no return
func bar(s string){
    fmt.Println("My name is ",s)
}

// one parameter, one return
func bazz(s string) string {
    return fmt.Sprint("They call me ", s)
}
fmt.Println(bazz("Mr. T"))
// [output] They call me Mr. T

// multiple parameters, multiple returns
func dogYears(name string, age int) (string, int){
    age *= 7
    return fmt.Sprint(name, " is this old in dog years "), age
}
s1, n := dogYears("Todd", 10)
fmt.Println(s1, n)
// [output] Todd is this old in dog years 70
```

## Variadic parameter

Can create a function that **takes an unlimited number of arguments**. This is known as a 'variadic parameter'. It **needs** to be the final parameter.

```golang
func sum(ii ...int){
    fmt.Println(ii)
    fmt.Printf("%T\n",ii)
}
sum(1,2,3,4,5,6,7,8,9)
/*
[output]
[1 2 3 4 5 6 7 8 9]
[]int
*/
```

## Deferring a function call

In Go, a 'defer' statement invokes a function whose **execution is deferred to the moment the surrounding function returns**, either because

- the surrounding function executed a **return** statement
- reached the **end** of its function body
- the corresponding goroutine is **panicking**

```golang
func main(){
    defer foo()
    bar()
}
func foo(){
    fmt.Println("foo")
}
func bar(){
    fmt.Println("bar")
}
// [output] bar\nfoo
```

This is mostly used for closing resources

```golang
import "os"
func main(){
f, err := os.Open("example.txt")
if err != nil {
    log.Fatalf("error: %s", err)
}
defer f.Close()
}
```

## Attaching methods to types

A method is a function attached to a type.

```golang
func (receiver ReceiverType) MethodName(parameters) ReturnType {

type person struct {
    first string
}
func (p person) speak(){
    fmt.Println("I am", p.first)
}

func main(){
    p1 := person{
        first: "James"
    }
    p1.speak()
    // [output] I am James
}
}
```

## Interfaces & polymorphism

An interface is a set of method signatures. It describes the behavior or contract that a type should adhere to. Any type that implements the methods specified in an interface is said to satisfy that interface. In Go, interfaces are implicitly implemented; there's no need to explicitly declare that a type implements an interface.

```golang
type Shape interface {
Area() float64
Perimeter() float64
}
type Circle struct {
radius float64
}
func (c Circle) Area() float64 {
return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
return 2 * math.Pi * c.radius
}
```

The following example above `^` implicitly implements the `Shape` interface.

## Anonymous functions

An anonymous function, often also called a function literal, lambda, or closure, is a function that is defined without being given a name.

```golang
// func (p parameter(s))(r return(s)) {...}
func(){
    fmt.Println("anon func ran")
}()
//anonymous function with parameters
func(s string){
    fmt.Println("This is my name ", s)
}("Todd")
```

## Function expressions

An expression is a combination of values, variables, operators, and function calls that are evaluated to produce a single value. It can be as simple as a literal value or a variable, or it can involve complex operations and function invocations.

```golang
x := func(){
    fmt.Println("Anon function assigned to x")
}
x()
```

## Returning a func

Returning a function in the Go programming language is a form of using higher-order functions, which are functions that can accept other functions as arguments and/or return other functions as results. This is a common practice in functional programming paradigms, but it's also available in multiparadigm languages like Go.

```golang
func main(){
    x := foo()
    fmt.Println(x)
    y := bar()
    fmt.Println(y())

    fmt.Printf("%T\n%T\n%T", foo, bar, y)
}
func foo() int {
    return 42
}
func bar() func() int {
    return func() int {
        return 43
    }
}
/*
    [output]
    42
    43
    foo() int
    func() func() int
    func() int
*/
```

### Callback

a function that is passed as an argument to another function

```golang
package main
import "fmt"
// Function that takes another function as an argument
func applyOperation(a, b int, operation func(int, int) int) int {
    return operation(a, b)
    }
// Functions to be used as arguments
func add(a, b int) int {
    return a + b
    }
func subtract(a, b int) int {
    return a - b
    }
func main() {
    result1 := applyOperation(5, 3, add) // Passing 'add' function as argument
    result2 := applyOperation(8, 4, subtract) // Passing 'subtract' function as argument
    fmt.Println(result1) // Output: 8
    fmt.Println(result2) // Output: 4
    }
```

## Closure

When one scope encloses over another scope. Variables that are declared in the outer scope are accessible in the inner scope. This helps limit the scope of variables

```golang
package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3
	fmt.Println(f()) // 4
	fmt.Println(f()) // 5
	fmt.Println(f()) // 6

	g := incrementor()
	fmt.Println(g()) // 1
	fmt.Println(g()) // 2
	fmt.Println(g()) // 3
	fmt.Println(g()) // 4
	fmt.Println(g()) // 5
	fmt.Println(g()) // 6
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
```

## Recursion

when a function calls itself.

```golang
func main() {
	// 4!
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(factLoop(4))

}

func factorial(n int) int {
	fmt.Println("This is n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

/*
factorial(4)		--> 4 *
factorial(4-1)		--> 3 *
factorial(3-1)		--> 2 *
factorial(2-1)		--> 1 *
factorial(1-1)		--> 0 // base case
*/

func factLoop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}
```
## Wrapper functions