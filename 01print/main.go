package main

import (
	"fmt"
)

func PrintFuncs() {

	const a, b = "test", 123

	fmt.Println(a, b)
	fmt.Print("Print")
	fmt.Println("Println")
	fmt.Printf("Verbs   :  %v, %T, %d, %.2f, %s, %q \n", 3, 3, 4, 4.555, "hello", "hello")
	fmt.Printf("%x %X %U \n", "a", "A", 'c')
}
