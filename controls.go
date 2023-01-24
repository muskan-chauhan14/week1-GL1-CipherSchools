package main

import "fmt"

/*
1. if-else statement
	if(condition) {
		//statements
	}

	if(condition){
		//statements
	} else {
		//statements
	}

	if(condition){
		//statements
	} else if {
		//statements
	} else {
		//statements
	}

	2. switch(data) {
	case var1:
		statements
	case var2:
		statements
	...
	default:
		statements
	}
*/

func main() {
	// fmt.Print("Enter a number: ")
	// var input int
	// fmt.Scanln(&input) //press enter as well

	// if input%2 == 0 {
	// 	fmt.Println("hey you are even")
	// } else {
	// 	fmt.Println("hey you are odd")
	// }

	// //multiple statements in golang
	// if x := 10; x%2 == 0 {
	// 	fmt.Println("hey you are even")
	// } else {
	// 	fmt.Println("hey you are odd")
	// }

	data := 100
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 103
		fmt.Println(data)
		fallthrough //executes the next case
	case 11:
		data = 104
		fmt.Println(data)
	case 15:
		data = 1002
		fmt.Println(data)
	default:
		data = 10909
		fmt.Println(data)
	}
}
