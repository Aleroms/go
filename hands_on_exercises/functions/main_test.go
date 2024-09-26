package main

import "testing"

func TestAdd63(t *testing.T) {
    total := Add63(5, 5)
    if total != 10 {
        t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}
func TestDoMath(t *testing.T){
    add := doMath(5,5,add)
    if add != 10 {
        t.Errorf("Sum was incorrect, got: %d, want: %d.", add, 10)
    }
    sub := doMath(10,5,subtract)
    if sub != 5 {
        t.Errorf("Subtraction was incorrect, got: %d, want: %d.",sub,5)
    }
}
func TestAdd64(t *testing.T){
    sum := add(40,2)
    if sum != 42 {
        t.Errorf("Subtraction was incorrect, got: %d, want: %d.",sum,42)
    }
}
func TestSubtract64(t *testing.T){
    sub := subtract(40,2)
    if sub != 38 {
        t.Errorf("Subtraction was incorrect, got: %d, want: %d.",sub,38)
    }
}
func TestBar58(t *testing.T){
    age, name := bar58()
    if age != 27 && name != "Santiago" {
        t.Errorf("Expected age:27, name:Santiago, Got age:%d, name:%s\n",age,name)
    }
}