package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 1, 10}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	fmt.Println(nums)
}
