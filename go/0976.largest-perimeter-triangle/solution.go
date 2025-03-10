// Created by lotus at 2025/03/10 22:34
// leetgo: 1.4.13
// https://leetcode.com/problems/largest-perimeter-triangle/

package main

import (
	"fmt"
	"sort"
)

// @lc code=begin

func largestPerimeter(nums []int) (ans int) {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	a, b, c := 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {
		c = nums[i]
		b = nums[i+1]
		a = nums[i+2]

		if a+b > c {
			return a + b + c
		}
	}

	return 0
}

// @lc code=end

func main() {
	nums := []int{3, 2, 3, 4}
	fmt.Println(largestPerimeter(nums))
}
