# Structs

## Introduction

A struct is a composite data type. (aggregate and complex data types). Structs allow us to compose together values of different types.

```golang
type person struct {
    first string
    last string
    age int
}
p1 := person {
    "James",
    "Bond",
    32
}
fmt.Printf("%T",p1)
// main.person
```

## Embedded Structs

We can take **one struct and embed it in another struct**. When you do this, the inner type gets promoted to the outer type

```golang
type secretAgent struct {
    person
    licenseToKill bool
}
p2 := secretAgent {
    person: person {
    "James",
    "Bond",
    32
}
    licenseToKill: true,
}
fmt.Printf("%T",p2)
// [output] main.secretAgent
// person stuff
fmt.Println(p2.first, p2.last, p2.age)
// secretAgent stuff
fmt.Println(p2.licenseToKill)
```

If both structs in the embedded struct have the same field, the outter-struct takes precidence.

```golang
type a struct{
    name string
}
type b struct {
    a
    name string
}
ex1:= a {
    name: "santi"
}
ex2: b {
    a: {name: "llama"},
    name: "meta"
}
fmt.Println(ex2.name)
// [output] meta
```
You can specify which.
```golang
fmt.Println(ex2.a.name)
// [output] llama
```
## Anonymous Structs
An anonymous struct is a struct which is not associated with a specific identifier.
```golang
p3 := struct {
    first string
    last string
    age int
}{
    first "Santiago",
    last "Morales",
    age: 27
}
fmt.Printf("%T",p3)
// [output] struct { first string; last string; age int}
```
## Composition Structs
Refers to the structuring/building of complex types by combining multiple simpler types. One way of composition is by *embedded structs*. The inner struct becomes accessible to the outer struct and is said to be **promoted** to the outer type. It is similar to inheritance in OOP but Go does *not* have iheritance.

*composite data types, aggregate data types, complex data types*

Composition is similar to inheritance such that the outer type can use the inner type's methods and fields. It creates a relationship between the types

A reason to create a type is to attach methods to it and self-document the code. For example, the [time](https://pkg.go.dev/time#Duration). Types are a way to encapsulate logic.

*anonymous structs can be inferred into named ones at compile time*