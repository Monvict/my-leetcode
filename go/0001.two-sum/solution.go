// Created by lotus at 2025/02/19 22:11
// leetgo: 1.4.13
// https://leetcode.com/problems/two-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func twoSum(nums []int, target int) []int {
	// key --> number
	// val --> index
	numMap := make(map[int]int)

	for i, num := range nums {

		if idx, ok := numMap[target-num]; ok {
			return []int{idx, i}
		}

		numMap[num] = i
	}

	return []int{999, 999}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := twoSum(nums, target)

	fmt.Println("\noutput:", Serialize(ans))
}
