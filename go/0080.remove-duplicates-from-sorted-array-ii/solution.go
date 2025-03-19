// Created by lotus at 2025/03/19 22:07
// leetgo: 1.4.13
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/

package main

import (
	"fmt"
)

// @lc code=begin

func removeDuplicates(nums []int) (ans int) {
	nunMaps := make(map[int]int, 0)

	pos := 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if nunMaps[num] <= 1 {
			nunMaps[num]++

			nums[pos] = num
			pos++
		}
	}

	fmt.Println(nums)
	return pos
}

// @lc code=end

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	removeDuplicates(nums)
}
