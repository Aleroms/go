package main

import "fmt"

type person struct {
	first_name         string
	last_name          string
	favorite_ice_cream []string
}

func main() {
	//create type person
	
	//create 2 values of type person

	// struct literal
	ex53a := person{
		first_name:         "Santiago",
		last_name:          "Morales",
		favorite_ice_cream: []string{"vanilla", "mint chocolate chip"},
	}
	ex53b := person{
		first_name:         "James",
		last_name:          "Moose",
		favorite_ice_cream: []string{"coconut"},
	}
	fmt.Printf("ex53 type is %T. Here are there contents...\n",ex53a)
	fmt.Println(ex53a)
	fmt.Println(ex53b.first_name + " " + ex53a.last_name)
	for _,v := range ex53b.favorite_ice_cream {
		fmt.Printf("\t%v\n",v)
	}
	
	/* 
	store the values of type person in a map with key = last name.
	access each. print out values, range over slice
	*/
	fmt.Println("------")
	ex54 := map[string]person {
		ex53a.last_name: ex53a,
		ex53b.last_name: ex53b,
	}
	fmt.Printf("ex54 type is %T. Here are there contents...\n",ex54)
	for key, value := range ex54 {
		fmt.Printf("%v %v likes...\n",ex54[key].first_name,ex54[key].last_name)
		for _, ice_cream := range value.favorite_ice_cream {
			fmt.Printf("\t%v\n",ice_cream)
		}
	}
	fmt.Println("------")
	/*
	create type engine and vehicle struct
	create two values of type vehicle using composite literal
	print out ea value and a single field from each of those values 
	*/
	type engine struct {
		electric bool
	}
	type vehicle struct {
		engine engine
		make string
		model string
		color string
		doors int
	}
	ex55a := vehicle{
		engine: engine{
			electric: false,
		},
		make: "Royal Enfield",
		model: "Continental GT 650",
		color: "white",
		doors: 0,
	}
	ex55b := vehicle {
		engine: engine{
			electric: true,
		},
		make: "Nissan",
		model: "Leaf",
		color: "grey",
		doors: 4,
	}
	fmt.Printf("ex55 type is %T. Here are there contents...\n",ex55a)
	ex55c := []vehicle{ex55a, ex55b}
	for _, v := range ex55c {
		var hasElectricEngine string
		
		if v.engine.electric {
			hasElectricEngine = "has an electric engine"
			}else {
				hasElectricEngine = "does not have an electric engine"
			}
			
			fmt.Printf("This %s %s is %s, %s and has %d doors.\n", v.make, v.model, v.color, hasElectricEngine, v.doors)
		}
		
		fmt.Println("------")
		// create an anonymous struct and print things
		ex56 := struct{
			first string
			friends map[string]int
			favDrinks []string
			}{
				first: "Santi",
				friends: map[string]int {
					"family": 10,
					"UCI": 13,
					"grade school": 5,
				},
				favDrinks: []string{"mango cart", "corona familiar", "brew X"},
			}
		fmt.Printf("ex56 type is %T. Here are there contents...\n",ex56)
		fmt.Println(ex56)

}