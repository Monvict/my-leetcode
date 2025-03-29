// Created by lotus at 2025/03/29 23:06
// leetgo: 1.4.13
// https://leetcode.com/problems/is-subsequence/

package main

import "fmt"

// @lc code=begin

func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for ; j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}

		if i >= len(s) {
			return true
		}
	}

	return false
}

// @lc code=end

func main() {
	s := "abc"
	t := "ahbgdc"

	fmt.Println(isSubsequence(s, t))
}
