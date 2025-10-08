package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruits = []string{"apple", "orange", "mango", "graphs"}
	fmt.Println(fruits)
	fmt.Println(fruits[1:])
	fmt.Println(fruits[:3])
	fmt.Println(fruits[1:3])
	fmt.Println(append(fruits, "Bannan"))

	numbers := make([]int, 4)

	numbers[0] = 200
	numbers[1] = 300
	numbers[2] = 100
	numbers[3] = 400
	// numbers[4] = 500 // error: index out of range [4] with length 4

	// but this works because reallocate the memory
	numbers = append(numbers, 500)
	fmt.Println(numbers)

	// Is Sorted
	fmt.Println("Is number sorted? ", sort.IntsAreSorted(numbers))

	// Sorting
	sort.Ints(numbers)
	fmt.Println(numbers)

	// Is Sorted
	fmt.Println("Is number sorted? ", sort.IntsAreSorted(numbers))

	// Remove the value from index

	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println(numbers)

}
