// Created by lotus at 2025/02/17 22:58
// leetgo: 1.4.13
// https://leetcode.com/problems/move-zeroes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func moveZeroes(nums []int)  {
    i := 0
	for j:=1; j < len(nums); j++ {

		// no-zero move to left
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]

			// incr left index 
			i++
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	moveZeroes(nums)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
