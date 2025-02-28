// Created by lotus at 2025/02/28 14:37
// leetgo: 1.4.13
// https://leetcode.com/problems/robot-return-to-origin/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func judgeCircle(moves string) bool {
	if len(moves) == 0 {
		return true
	}

	x := 0
	y := 0

	for _, m := range moves {
		switch m {
		case 'L':
			x -= 1
		case 'R':
			x += 1
		case 'U':
			y += 1
		case 'D':
			y -= 1
		}
	}

	return x == 0 && y == 0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	moves := Deserialize[string](ReadLine(stdin))
	ans := judgeCircle(moves)

	fmt.Println("\noutput:", Serialize(ans))
}
