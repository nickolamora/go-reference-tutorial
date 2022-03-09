package main

import "fmt"

// map is a data type of <K,V>

func main() {
	shoppingCart := make(map[string]int)
	shoppingCart = map[string]int{
		"keyboard": 100,
		"mouse":    50,
		"laptop":   1000,
	}
	fmt.Println(shoppingCart)
	fmt.Println(len(shoppingCart))
	shoppingCart["keyboard"] = 250
	fmt.Println(shoppingCart["keyboard"])
	shoppingCart["cpu"] = 400
	fmt.Println(shoppingCart)

	//checking if there is an item in a map. If it doesnt exist, returns false
	//cart, ok := shoppingCart["phone"]
	//fmt.Println(cart, ok)
	// or so that you dont have to use a value use _
	_, ok := shoppingCart["something"]
	fmt.Println(ok)

	//when copying data from one map to another it copies the ADDRESS not the value similar to slices
	sc := shoppingCart
	sc["cpu"] = 10000
	fmt.Println(shoppingCart)
	fmt.Println(sc)

	//removing data
	delete(sc, "cpu")
	fmt.Println(sc)
	fmt.Println(shoppingCart)

}
