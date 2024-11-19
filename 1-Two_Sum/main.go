package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	rmap := make(map[int]int)

	for i, num := range nums {
		need := target - num
		if index, exists := rmap[need]; exists {
			return []int{index, i}
		}

		rmap[num] = i
	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
