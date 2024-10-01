# Pointers

A pointer refers to a **variable that holds the memory address**. Pointers allow you to directly manipulate memory and build data structures.

- Address operator `&`
  - gets the memory address of a variable.
- Defererencing operator `*`
  - gets the value stored at the memory address.

```golang
func main(){
    i:=42
    ptr := &i
    fmt.Println(*ptr)
    fmt.Printf("Type of ptr: %T\n", ptr)
    *ptr = 21
    fmt.Println(i)
}
/*
    output
    42
    Type of ptr: *int
    21
*/
```

## Pass by value/ reference types, and mutability

Go has several reference types, including _pointers, slices, maps, channels, functions and interfaces_

1. **Pointers**: A pointer holds the memory address of a value. It allows you to directly access and
   modify the memory location of a value.
2. **Slices**: A slice is a descriptor of an array segment. It includes a pointer to the array, the
   length of the segment, and its capacity (the maximum length of the segment).
3. **Maps**: A map is a powerful data structure that associates values of one type (the key) with
   values of another type (the value). It's an unordered collection of key-value pairs.
4. **Channels**: Channels are used for communication between goroutines (Go's term for
   threads). They allow you to pass data from one goroutine to another.
5. **Functions**: In Go, functions are first-class citizens, meaning they can be assigned to
   variables, passed as arguments to other functions, and returned as values from other functions.
   When a function is assigned to a variable, the variable stores a reference to the function.
6. **Interfaces**: An interface type represents a set of method signatures. It provides a way to
   specify the behavior of an object: if something can do this, then it can be used here.

remember, Go passes **by value**. Any time immutable data types like int, float, string, are passed to a function, the changes to the variable will not affect the original value.

In summary, in Go, everything is passed by value. However, for mutable data types, the
"value" that's passed includes a reference to the underlying data, which is why changes
made inside the function are visible outside of it.

## Pointer and value heuristics

Values should be default. Pointer semantics have advantages in some situations but opens up for complexity.

- use **value semantics** when possible
  - they are simpler and safer since they don't share state or require memory management.
  - if a `func` does not need to modify the input, or the data is small, use value semantics
- use **pointer semantics** for large data
  - copying large structs/arrays can be inefficient
- use pointer semantics for **mutability**
  - if the `func` _needs_ to modify its receiver/input parameter
  - common for methods that need to update state in struct
- stay consistant
  - use either pointer semantics or value semantics and stick to it
- pointer semantics when **interfacing** with other code

## Pointer, values, the stack & the heap

When a function receives a value that isn't a pointer, it gets its _own copy_ of that value and it is placed on the **stack**. This is fast and doesn't invluve any garbage collection. One the function returns, this memory is reclaimed.

For **pointer**, when a function receives a pointer or returns a pointer, it indicates to the compiler that could be shared across goroutine boundaries. The compiler **must** allocate it on the **heap**. Heap allocation is more expensive and requires garbage collection.

In Go, **escape analysis** decides whether a variable should be allocated on te stack or the heap. It examines the function to see if the pointer is being returned or if the variable is within a function literal (closure). If so, the variable escapes to the heap.

to see escape analysis, use `go run -gcflags -m main.go`

## Method sets

In Go, a `method set` is the _set of methods attached to a type_. This is key to Go's interface mechanism. It is associated with both value and pointer types.

- The method set of a **type T** consists of all methods with **receiver type T**.
  - can be called using variable of _type T_
- The method set of a **type \*T** consists of all methods with **receiver \*T or T**
  - can be called using variables of type pointer **\*T** as well as non-pointer types

```golang
type T struct {...}
func (t T) M1(){...}
func (t *T) M2(){...}
type X interface{
   M1()
   M2()
}
func main(){
   var n X
   t := T{}
   n = t // compoilation error, because T does not implement M2

   tPtr := &T{}
   n = tPtr // this is fine, *T implements M1 and M2
}
```
