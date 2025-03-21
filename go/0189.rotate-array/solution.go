// Created by lotus at 2025/03/21 17:55
// leetgo: 1.4.13
// https://leetcode.com/problems/rotate-array/

package main

import (
	"fmt"
)

// @lc code=begin
// this solution is so-called three times rotate method
// because (X^TY^T)^T = YX
func rotate(nums []int, k int) {
	k = k % len(nums)

	// 1. reverse first half
	reverse(nums, 0, len(nums)-k-1)

	// 2. revers second half
	reverse(nums, len(nums)-k, len(nums)-1)

	// 3. reverse the whole array
	reverse(nums, 0, len(nums)-1)

	fmt.Println(nums)
}

func reverse(nums []int, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func rotate2(nums []int, k int) {

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
	nus := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nus, 3)
	fmt.Println(nus)
}
