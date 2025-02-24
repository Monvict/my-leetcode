// Created by lotus at 2025/02/24 22:10
// leetgo: 1.4.13
// https://leetcode.com/problems/plus-one/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func plusOne(digits []int) (ans []int) {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	// the fist one is zero, means we need append 1 in the start
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	digits := Deserialize[[]int](ReadLine(stdin))
	ans := plusOne(digits)

	fmt.Println("\noutput:", Serialize(ans))
}
