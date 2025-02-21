// Created by lotus at 2025/02/21 22:16
// leetgo: 1.4.13
// https://leetcode.com/problems/find-the-difference/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findTheDifference(s string, t string) (ans byte) {
	cmap := make(map[rune]int, 0)

	for _, r := range s {
		cmap[r]++
	}

	for _, r := range t {
		cmap[r]--

		if cmap[r] == 0 {
			delete(cmap, r)
		}
	}

	var resut byte
	for c := range cmap {
		resut = byte(c)
	}

	return resut
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := findTheDifference(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
