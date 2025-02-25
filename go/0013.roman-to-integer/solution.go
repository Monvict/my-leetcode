// Created by lotus at 2025/02/25 22:55
// leetgo: 1.4.13
// https://leetcode.com/problems/roman-to-integer/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func romanToInt(s string) (ans int) {
	if len(s) == 0 {
		return 0
	}

	numMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	for i := 0; i < len(s)-1; i++ {
		// 每次比较2个字符
		// 当前一个 >= 后一个；则直接将前一个字符对应的数字加到resut
		// 否则：减去前一个数字(因为下一次循环会加上这一步的后一个)
		first := numMap[rune(s[i])]
		second := numMap[rune(s[i+1])]

		if first >= second {
			result += first
		} else {
			result -= first
		}

		// 将最后一个字符加上
		if i == len(s)-2 {
			result += second
		}
	}
	return result
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := romanToInt(s)

	fmt.Println("\noutput:", Serialize(ans))
}
