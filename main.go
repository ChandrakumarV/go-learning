package main

import (
	"fmt"
	"go-learning/hello"
	"os"
)

func main() {
	const a, b = "test", 123
	fmt.Println(a, b)

	// cmd args
	fmt.Println("Args Len: ", len(os.Args))

	// module import
	a1 := hello.Sum1(1, 2)
	a2 := hello.Sum2(3, 4)

	fmt.Println(a1)
	fmt.Println(a2)

	//main package helper (go run .)
	fmt.Println(Sum(10, 20))

}
