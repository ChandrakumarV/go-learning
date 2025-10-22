package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}

func main() {

	arr := []int{3, 6, 4, 2, 7, 8, 2, 11}

	result := containsDuplicate(arr)
	fmt.Println(result)
}
