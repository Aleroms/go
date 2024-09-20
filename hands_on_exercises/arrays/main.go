package main

import "fmt"

func main() {
	//composite literal
	arr := [5]int{}
	//assign 5 values to each index
	for i := 0; i < 5; i++ {
		arr[i] = i
	}
	//print the values and type using range
	for i := range arr {
		fmt.Printf("%v %T ",i,i)
	}

	//composite literal slice
	slice := []int{42,43,44,45,46,47,48,49,50,51}

	for i,v := range slice{
		fmt.Printf("%v %T %v ",v,v,i)
	}
	//[inclusive:exclusive]
	fmt.Println(slice[:5])
	fmt.Println(slice[5:])
	fmt.Println(slice[2:7])
	fmt.Println(slice[1:6])
	//create new slices using slice

	x := []int{42,43,44,45,46,47,48,49,50,51}
	x = append(x,52)
	fmt.Println(x)
	x = append(x,53,54,55)
	fmt.Println(x)
	y := []int{56,57,58,59,60}
	x = append(x, y...)
	fmt.Println(x)
	x = []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(x)
	x = append(x[0:3],x[6:]...)
	fmt.Println(x)

	usa_states := make([]string,50)
	usa_states = append(usa_states, "alaska","arizona","arkansa","california","conneticut","delaware","florida","georgia","illinois", "indiana", "iowa")
	fmt.Println(len(usa_states))
	fmt.Println(cap(usa_states))

	jb := []string{"james","bond", "shaken, not stirred"}
	jm := []string{"miss", "moneypenny", "i'm 008"}
	multidemensional := [][]string{jb,jm}
	for i,v := range multidemensional{
		fmt.Println(i,v)
		for a, b := range v {
			fmt.Println(a,b)
		}
	}

	
}