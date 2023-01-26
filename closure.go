package main

//closure: function passed as a prameter

import "fmt"

func main() {

	//Function is a first class member in golang
	//number:=10
	var number int
	number = 10
	fmt.Println(number)

	//store a function as a vale in a variable
	//var variable_name datatype
	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In a function")
		return 10 + x
	}
	j := getInt(19)
	i := func(x int) int {
		fmt.Println("in a function")
		return 20 + x
	}(19) //Anonymous function: declaration and calling of function; NO NAME
	fmt.Println(i)
	fmt.Println(j)

	var y func()
	y = func() {
		fmt.Print(9)
	}
	y()

	g(getInt)
}

func g(getInt func(int) int) {
	getInt(78)
}
