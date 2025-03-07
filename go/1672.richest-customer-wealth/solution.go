// Created by lotus at 2025/03/07 08:42
// leetgo: 1.4.13
// https://leetcode.com/problems/richest-customer-wealth/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumWealth(accounts [][]int) (ans int) {

	max := 0

	for _, account := range accounts {
		s := sumArray(account)
		if s >= max {
			max = s
		}
	}
	return max
}

func sumArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	sum := 0
	for _, a := range arr {
		sum += a
	}

	return sum
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	accounts := Deserialize[[][]int](ReadLine(stdin))
	ans := maximumWealth(accounts)

	fmt.Println("\noutput:", Serialize(ans))
}
