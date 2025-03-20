// Created by lotus at 2025/03/20 23:24
// leetgo: 1.4.13
// https://leetcode.com/problems/majority-element/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func majorityElement(nums []int) (ans int) {
	numMap := make(map[int]int)
	len := len(nums)
	for _, num := range nums {
		numMap[num]++

		if numMap[num] > len/2 {
			return num
		}
	}
	return 0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := majorityElement(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
