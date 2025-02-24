// Created by lotus at 2025/02/24 22:37
// leetgo: 1.4.13
// https://leetcode.com/problems/sign-of-the-product-of-an-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func arraySign(nums []int) (ans int) {
	negCnt := 0

	for _, num := range nums {
		if num == 0 {
			return 0
		}

		if num < 0 {
			negCnt++
		}
	}

	if negCnt%2 == 0 {
		return 1
	}
	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := arraySign(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
