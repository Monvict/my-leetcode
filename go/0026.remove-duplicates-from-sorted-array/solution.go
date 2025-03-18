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
	numMap := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if numMap[num] == 0 {
			numMap[num]++

			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}

// @lc code=end

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("pos = ", removeDuplicates(nums))
}
