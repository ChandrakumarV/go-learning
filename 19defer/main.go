package main

import "fmt"

// defer is call later, usually for purposes of cleanup
// deferred calls are executed in LIFO order

func main() {

	defer fmt.Println("world")
	fmt.Println("hello")

	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
