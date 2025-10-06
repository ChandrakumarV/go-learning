package main

import "fmt"

func main() {

	// Method 1 arr declaration ----------
	var arr [3]int
	fmt.Println(arr)
	arr[0] = 100
	arr[1] = 200
	arr[2] = 300
	fmt.Printf("Type of array is %T \n", arr)
	fmt.Println("Numbers are arr : ", arr)
	fmt.Println("Lenght of No : ", len(arr))

	// Method 2 arr declaration ----------
	var fruits = []string{"apple", "orange", "mango", "graphs"}
	fmt.Println(fruits)

}
