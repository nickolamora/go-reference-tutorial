package main

import "fmt"

func main() {

	//if
	if true {
		fmt.Println("Hi")
	} else {
		fmt.Println("Bye")
	}

	//switch, no breaks statements needed. If you need to control it, however, you can use 'break'
	// fallthrough can also be used to execute the statement right after it ( not allowed in type switching)
	switch 2 {
	case 1, 3, 5:
		fmt.Println("this is 1")
		break
	case 2, 4, 6:
		fmt.Println("this is 2")
		fallthrough
	default:
		fmt.Println("default")
	}

	//type switching, doing something based on the type
	var i interface{} = 5
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("default")
	}
}
