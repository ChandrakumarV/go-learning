package main

import "fmt"

func main() {
	arr := []string{"chandru", "naresh", "giri", "surya"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for idx, val := range arr {
		fmt.Println(idx, " : ", val)
	}

	for _, val := range arr {
		fmt.Println(val)
	}

	// Continue
	for i := 1; i <= 10; i++ {
		if i == 5 {
			fmt.Println("-------")
			continue
		}
		fmt.Println(i)
	}

	// Break
	for i := 1; i <= 10; i++ {
		if i == 8 {
			fmt.Println("*******")
			break
		}
		fmt.Println(i)
	}

	// Goto
	for i := 11; i <= 20; i++ {
		if i == 15 {
			goto hello
		}
		fmt.Println(i)
	}

hello:
	fmt.Println("GOTO")
}
