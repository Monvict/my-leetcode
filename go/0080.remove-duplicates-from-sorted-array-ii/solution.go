// Created by lotus at 2025/03/19 22:07
// leetgo: 1.4.13
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

package main

import (
	"fmt"
)

// @lc code=begin

func removeDuplicates(nums []int) (ans int) {

	pos := 0
	repeat := 1

	for i := 1; i < len(nums); i++ {
		if nums[pos] == nums[i] {
			repeat++

			if repeat <= 2 {
				nums[pos+1] = nums[i]
				pos++
			}
		} else {
			nums[pos+1] = nums[i]
			pos++
			repeat = 1
		}

	}

	fmt.Println(nums)
	return pos + 1
}

// @lc code=end

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println("pos", removeDuplicates(nums))
}
