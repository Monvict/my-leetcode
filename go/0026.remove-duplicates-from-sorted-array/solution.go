// Created by lotus at 2025/03/18 23:31
// leetgo: 1.4.13
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package main

import (
	"fmt"
)

// @lc code=begin

func removeDuplicates(nums []int) (ans int) {
	pos := 0

	// i = 1 --> means skip the fisrt one
	for i := 1; i < len(nums); i++ {
		if nums[pos] != nums[i] {
			nums[pos+1] = nums[i]
			pos++
		}
	}

	fmt.Println(nums)
	return pos + 1
}

// @lc code=end

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("pos = ", removeDuplicates(nums))
}
