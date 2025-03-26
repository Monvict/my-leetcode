// Created by lotus at 2025/03/26 20:23
// leetgo: 1.4.13
// https://leetcode.com/problems/find-the-duplicate-number/

package main

import (
	"fmt"
)

// @lc code=begin

func findDuplicate(nums []int) (ans int) {
	low := 0
	hight := len(nums) - 1

	for low < hight {
		cnt := 0
		mid := (low + hight) / 2
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}

		// duplicate num in left side of mid
		if cnt > mid {
			hight = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

// @lc code=end

func main() {
	nums := []int{1, 2, 4, 3, 2}
	fmt.Println(findDuplicate(nums))
}
