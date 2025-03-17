// Created by lotus at 2025/03/17 22:39
// leetgo: 1.4.13
// https://leetcode.com/problems/remove-element/

package main

import (
	"fmt"
)

// @lc code=begin

func removeElement(nums []int, val int) (ans int) {
	i, j := len(nums)-1, len(nums)-1
	for ; j >= 0; j-- {
		if nums[j] != val {
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		i--
	}

	fmt.Println(nums)
	fmt.Printf("i = %d, j = %d\n", i, j)
	return i + 1
}

// @lc code=end

func main() {
	nums := []int{3, 2, 2, 3}
	removeElement(nums, 3)
}
