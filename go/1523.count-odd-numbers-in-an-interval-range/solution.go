// Created by lotus at 2025/03/10 20:04
// leetgo: 1.4.13
// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func countOdds(low int, high int) (ans int) {

	for i := low; i <= high; i++ {
		if i%2 != 0 {
			ans++
		}
	}
	return ans
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	low := Deserialize[int](ReadLine(stdin))
	high := Deserialize[int](ReadLine(stdin))
	ans := countOdds(low, high)

	fmt.Println("\noutput:", Serialize(ans))
}
