package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// Check if length 0 or 1
	if len(nums) <= 1 {
		return false
	}

	// create hashmap for O(n)
	have := map[int]struct{}{}

	for _, v := range nums {
		_, exists := have[v]
		if exists {
			return true
		}
		have[v] = struct{}{}
	}

	// No duplicates = false
	return false
}

func main() {
	data := []int{1, 2, 3}
	fmt.Println(containsDuplicate(data))
}
