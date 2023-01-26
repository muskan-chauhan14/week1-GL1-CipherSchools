package main

import "fmt"

/*
func function_name() {
	statement-1
	statement-2
	statement-3
	statement-4
}
*/

func main() {
	w := new(int) //pointer
	name := new(string)
	res := a(w, name) //12
	fmt.Println(res)  //muskan

	t := 23
	r1, r2 := b(&t, *w)      //pointer
	fmt.Println(r1, " ", r2) //122 19

	d(1, 2, 3, 4, 5, 6)

	fmt.Println(name)  //address
	fmt.Println(*name) //value

}

func a(w *int, name *string) (x string) {
	fmt.Println(12)
	x = "muskan"
	*w = 100
	*name = "muskan"
	return
}

func b(*int, int) (int, int) {
	return c()
}
func c() (int, int) {
	return 122, 19
}

func d(args ...int) {
	fmt.Println(args) //list as output; can be used for range
	//[1 2 3 4 5 6]
	for _, v := range args {
		fmt.Print(v) //123456
	}
}
