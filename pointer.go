package main

import "fmt"

func main() {
	// i := 10
	// j := &i
	// *j = 100

	var i int
	i = 10
	//var j *int //declration, j is empty
	//<nil>
	j := new(int) //declaration + assignment, j is not empty
	//random address
	j = &i
	*j = 100

	fmt.Println(*j)
	fmt.Println(i)

	name := new(string)
	*name = "golang"
	fmt.Println(*name)
}
