// Created by lotus at 2025/03/10 20:14
// leetgo: 1.4.13
// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func average(salary []int) (ans float64) {
	min, max, sum := salary[0], salary[0], 0

	for i := 0; i < len(salary); i++ {
		cur := salary[i]
		sum += cur

		if cur <= min {
			min = cur
		}

		if cur >= max {
			max = cur
		}
	}

	ans = float64((sum - min - max)) / float64((len(salary) - 2))
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	salary := Deserialize[[]int](ReadLine(stdin))
	ans := average(salary)

	fmt.Println("\noutput:", Serialize(ans))
}
