package main

import "fmt"

// find the sum two no = 19 and return the indices

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)

	for i, num := range nums {
		neededNo := target - num

		if mapIndex, isContain := m[neededNo]; isContain {
			return []int{mapIndex, i}
		}

		m[num] = i

	}

	return []int{}

}

func main() {

	tar := 19
	arr := []int{3, 6, 4, 7, 8, 2, 11}

	result := twoSum(arr, tar)
	fmt.Println(result)
}
