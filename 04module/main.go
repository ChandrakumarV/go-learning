package main

import (
	"04module/hello"
	"fmt"
)

func main() {

	// From diff package
	a1 := hello.Sum1(1, 2)
	a2 := hello.Sum2(3, 4)

	// From main package itself
	Helper()

	fmt.Println("Sum 1 :", a1)
	fmt.Println("Sum 2 :", a2)

}
