// Created by lotus at 2025/03/26 23:27
// leetgo: 1.4.13
// https://leetcode.com/problems/valid-palindrome/

package main

import (
	"fmt"
	"strings"
	"unicode"
)

// @lc code=begin

func isPalindrome(s string) bool {
	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	s = strings.Map(f, s)

	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

// @lc code=end

func main() {
	s := "A man, a plan, a canal: Panamax"
	fmt.Println(isPalindrome(s))
}
