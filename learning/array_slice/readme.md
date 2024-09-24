# Introduction to grouping values

## array
array - a numbered sequence of **VALUES** of the same **TYPE**
- it does not change in size; specified at init
- used for GO internales; generally not used in code
resource - [documentation](https://golang.org/ref/spec#Array_types)

## slice
slice - is built on top of an array
- dynamic size
- holds the **VALUES** of the same **TYPE**
- has a length and a capacity property
[documentation](https://golang.org/ref/spec#Slice_types)

## map
map - is a **key/value** storage
- an *unordered* group of VALUES of one TYPE, 
- a set of unique keys of another type
[documentation](https://golang.org/ref/spec#Map_types)

## struct
struct - a data structure
- a composite / aggregate
- allows us to **collect values of different types together**
[documentation](https://golang.org/ref/spec#struct_types)
