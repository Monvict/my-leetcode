// Created by lotus at 2025/02/22 22:22
// leetgo: 1.4.13
// https://leetcode.com/problems/valid-anagram/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isAnagram(s string, t string) bool {
	cmap := make(map[rune]int)

	for _, c := range s {
		cmap[c]++
	}

	for _, c := range t {
		cmap[c]--
	}

	for _, v := range cmap {
		if v > 0 {
			return false
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := isAnagram(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
