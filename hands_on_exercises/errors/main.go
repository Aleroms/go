package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

type santiError struct {
	s string
}

func (s santiError) Error() string {
	return s.s
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("JSON did not marshal, " ,err)
	}
	fmt.Println(string(bs))

	// testing is type error
	santiTestError := santiError {
		s : "test santi",
	}
	testIsError(santiTestError)

}
// toJSON needs to return an error also
func toJSON(a any) ([]byte, error) {
	bs, err := json.Marshal(a)
	if(err != nil){
		// return []byte{}, errors.New("cannot marshal")
		return []byte{}, fmt.Errorf("cannot marshal")
	}
	return bs, nil
}

func testIsError(err error){
	// calls the method set Error bcz it to is an error
	fmt.Println(err.Error())
}