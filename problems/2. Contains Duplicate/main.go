package main

import "fmt"

func containsDuplicateHashMap(nums []int) bool {
	m := make(map[int]bool)

	for _, num := range nums {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}

func containsDuplicateHashSet(nums []int) bool {
	s := make(map[int]struct{})

	for _, num := range nums {
		if _, found := s[num]; found {
			return true
		}
		s[num] = struct{}{}
	}
	return false
}

func main() {

	arr := []int{3, 6, 4, 2, 7, 8, 2, 11}

	resultHashMap := containsDuplicateHashMap(arr)
	resultHashSet := containsDuplicateHashSet(arr)

	fmt.Println("Hash Map:", resultHashMap)
	fmt.Println("Hash Set:", resultHashSet)
}
