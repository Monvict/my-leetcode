// Created by lotus at 2025/03/21 17:55
// leetgo: 1.4.13
// https://leetcode.com/problems/rotate-array/

package main

import (
	"fmt"
)

// @lc code=begin

func rotate(nums []int, k int) {

	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		shiftOneStep(nums)
		nums[0] = last
	}

	fmt.Println(nums)
}

func shiftOneStep(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		nums[i+1] = nums[i]
	}
}

// @lc code=end

func main() {
	nus := []int{-1, -100, 3, 99}
	rotate(nus, 2)
}
