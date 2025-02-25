// Created by lotus at 2025/02/25 22:32
// leetgo: 1.4.13
// https://leetcode.com/problems/monotonic-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isMonotonic(nums []int) bool {
	if len(nums) < 2 {
		return true
	}

	monotonicUp := true
	monotonicDown := true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			monotonicUp = false
		}

		if nums[i+1] > nums[i] {
			monotonicDown = false
		}
	}

	return monotonicDown || monotonicUp
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := isMonotonic(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
