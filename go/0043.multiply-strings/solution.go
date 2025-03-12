// Created by lotus at 2025/03/12 10:35
// leetgo: 1.4.13
// https://leetcode.com/problems/multiply-strings/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func multiply(num1 string, num2 string) string {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num1 := Deserialize[string](ReadLine(stdin))
	num2 := Deserialize[string](ReadLine(stdin))
	ans := multiply(num1, num2)

	fmt.Println("\noutput:", Serialize(ans))
}
