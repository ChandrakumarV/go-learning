package main

import "fmt"

// find the sum two no = 19 and return the indices

func main() {

	tar := 19
	arr := []int{3, 6, 4, 7, 8, 2, 11}

	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {

		neededNo := tar - arr[i]
		mapIndex, isContain := m[neededNo]

		if isContain {
			fmt.Println(mapIndex, i)
			return
		}

		m[arr[i]] = i

	}

	fmt.Println([]int{})
}
