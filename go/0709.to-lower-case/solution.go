// Created by lotus at 2025/02/19 23:17
// leetgo: 1.4.13
// https://leetcode.com/problems/to-lower-case/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func toLowerCase(s string) string {
	r := []byte{}
	for _, c := range s {
		charByte := byte(c)
		fmt.Println("charByte = ", charByte)
		if charByte >= 65 && charByte <= 90 {
			charByte = charByte + 32
		}

		r = append(r, charByte)
	}

	return string(r)
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := toLowerCase(s)

	fmt.Println("\noutput:", Serialize(ans))
}
