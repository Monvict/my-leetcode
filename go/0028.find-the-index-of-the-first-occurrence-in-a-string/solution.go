// Created by lotus at 2025/02/21 22:52
// leetgo: 1.4.13
// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func strStr(haystack string, needle string) (ans int) {

	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle); i++ {
		s := haystack[i : len(needle)+i]

		if s == needle {
			return i
		}
	}
	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	haystack := Deserialize[string](ReadLine(stdin))
	needle := Deserialize[string](ReadLine(stdin))
	ans := strStr(haystack, needle)

	fmt.Println("\noutput:", Serialize(ans))
}
