// Created by lotus at 2025/02/26 22:45
// leetgo: 1.4.13
// https://leetcode.com/problems/length-of-last-word/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLastWord(s string) (ans int) {
	if len(s) == 0 {
		return 0
	}

	r := len(s) - 1

	// right edge of last word
	for s[r] == ' ' {
		r--
	}

	// left edge of last word
	l := r
	for l > 0 && s[l] != ' ' {
		l--
	}
	return r - l
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLastWord(s)

	fmt.Println("\noutput:", Serialize(ans))
}
