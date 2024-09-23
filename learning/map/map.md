# Map

## Introduction

A map is a **key-value** store. The syntax for a map is

```golang
map[key]value
```

The key can be of any type which allows comparison. It is important to note that _maps are unordered_. They will print out in random order.

A map is a perfect data structure that has fast acceess

```golang
m := map[string]int {
  "todd": 42,
  "henry": 16,
  }

// or

m := make(map[string]int)
m["todd"] = 42
m["henry"] = 16

fmt.Println("Henry's age is ",m["henry"])
```

### Inserting an element in map

```golang
m["todd"] = 42
m["henry"] = 16
```

### Mapping over a map

This is how you map over a map

```golang
for k := range m {
  fmt.Printf("just the keys: %s\n",k)
}
```

### Deleting an element

This is how you delete a element from a map

```golang
ex := map[string]int{
  "ex1":1,
  "ex2":2,
  }
// deleting...
delete(ex, "ex1")
```

you delete by key name. If it does not exist, Golang will not panic. It will return 0. You can use the comma ok idiom to check if you have an existing key or not.

```golang
an := map[string]int {
  "Henry":1,
  "Tovs": 2,
  "Elizabeth":2
}
v, ok := an["Georgey"]
if ok {
  fmt.Println(v)
} else {
  fmt.Println("key dne")}
}
```

Because `Georgey` DNE, the following code will print `Key DNE`. We can use the comma ok idiom with the statement-statement idiom to condense the code down.

```golang
an := map[string]int {
  "Henry":1,
  "Tovs":2,
  "Elizabeth":2
}
if v, ok:= an["Lucas"]; !ok {
  fmt.Println("key dne")
} else {
  fmt.Println(v)
}
```

This is much more condensed as well as better because it limits the scope of the `ok` variable to only that condition.
