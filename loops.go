package main

import "fmt"

func main() {
	i := 0

	for i < 3 {
		if i == 1 {
			continue
		}
		fmt.Println(i)
		break
	}
	//o/p: 0

	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Println(i)
		//break
	}
	//o/p: 0
	//     2

	//list:=range(2)  //list cannot be used individually
	nums := []int{1, 3, 2, 4, 0}
	for index, val := range nums {
		fmt.Print(index, " ")
		fmt.Println(val)
	}
	/*o/p: 0 1
	       1 3
	       2 2
		   3 4
		   4 0
	*/

	for _, val := range nums { //range returns two values: index and the element
		if val == 3 {
			continue
		} // 1 2 4 0

		/*
			if val==3 {
				break
			} // 1  */

		fmt.Println(val)
	}
	/*o/p: 0 1
	       1 3
	       2 2
		   3 4
		   4 0
	*/

	for _, value := range "muskan" {
		fmt.Println(value) //ASCII values
	}

}
