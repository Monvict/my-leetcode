// Created by lotus at 2025/02/26 23:02
// leetgo: 1.4.13
// https://leetcode.com/problems/baseball-game/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func calPoints(operations []string) (ans int) {
	nums := []int{}

	for _, opr := range operations {
		n := len(nums)
		switch opr {
		case "+":
			nums = append(nums, nums[n-1]+nums[n-2])
		case "D":
			nums = append(nums, nums[n-1]*2)
		case "C":
			nums = nums[:n-1]
		default:
			r, _ := strconv.Atoi(opr)
			nums = append(nums, r)
		}
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	operations := Deserialize[[]string](ReadLine(stdin))
	ans := calPoints(operations)

	fmt.Println("\noutput:", Serialize(ans))
}
