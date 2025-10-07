package main

import "fmt"

func main() {

	// Method 1
	y := map[string]string{"name": "Chandru", "course": "Go"}
	fmt.Println(y)

	// Method 2
	x := make(map[string]int)
	x["age"] = 30
	x["year"] = 2023

	fmt.Println(x)
	fmt.Println(x["age"])

	// Looping through map
	for k, v := range x {
		fmt.Println(k, " : ", v)
	}

	// Checking key exists
	v, ok := x["age"]
	fmt.Println(v, ok)

	// Deleting key
	delete(x, "year")
	fmt.Println(x)

	// Checking key exists
	v2, ok2 := x["year"]
	fmt.Println(v2, ok2)
}
