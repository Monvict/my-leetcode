// Created by lotus at 2025/02/24 22:43
// leetgo: 1.4.13
// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}

	sort.Ints(arr)

	diff := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		span := arr[i+1] - arr[i]

		if span != diff {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[[]int](ReadLine(stdin))
	ans := canMakeArithmeticProgression(arr)

	fmt.Println("\noutput:", Serialize(ans))
}
