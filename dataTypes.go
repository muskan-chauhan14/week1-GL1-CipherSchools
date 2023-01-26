package main

import "fmt"

func main() {

	var data int
	//declaration
	//var Variable_name datatype

	data = 10 //initialization
	fmt.Println(data)

	data1 := 100
	fmt.Println(data1)

	//data:=1000 //redeclare the variable  ERROR
	data = 1000 //reassigning value

	//datatype
	//1. Primitive:
	//int float bool string complex byte
	//2. Non-Primitive
	//struct map chan interface array slice

	//functional programming language

	name := "rahul"
	fmt.Println(name)
	
	//const value must be initialized during declaration
	//const l int
	//l = 10

	const pi=3.14   //implicit typing
	const val int32 = 3 //explicit typing
	fmt.Println(pi)
}
