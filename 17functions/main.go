package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	printMessage()

	result := add2(3, 5)
	fmt.Println("The sum of 3 and 5 is:", result)

	// Multiple parameter
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("The total sum is:", total)

	// Multiple return values
	name, age := getUserInfo()
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// We can't define a function inside another function in Go.

	// func printMessageInner() {
	// 	fmt.Println("Hello from printMessageInner!")
	// }

}

func printMessage() {
	fmt.Println("Hello from printMessage!")
}

func add2(a int, b int) int {
	return a + b
}

func sumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getUserInfo() (string, int) {
	return "chandru", 32
}
