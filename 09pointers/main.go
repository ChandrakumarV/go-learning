package main

import "fmt"

func main() {

	// Method 1 (var) -----------
	var p1 *int

	fmt.Println(p1) // nil

	p1 = new(int)

	fmt.Println(p1)  // address of memory
	fmt.Println(&p1) // address of p1 variable
	fmt.Println(*p1) // 0

	// Method 2 (:) --------------
	p2 := new(int)
	fmt.Println(p2) //0

	// Reference -----------
	var num = 4343
	var ref = &num
	*ref = 4000
	*ref *= 2
	fmt.Println(num) //8000

}
