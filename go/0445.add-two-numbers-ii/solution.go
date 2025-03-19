// Created by lotus at 2025/03/17 22:29
// leetgo: 1.4.13
// https://leetcode.com/problems/add-two-numbers-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ans *ListNode) {
	// 1. 反转链表 #206

	// 2. 链表相加 #2

	// 3. 得到结果
	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	l1 := Deserialize[*ListNode](ReadLine(stdin))
	l2 := Deserialize[*ListNode](ReadLine(stdin))
	ans := addTwoNumbers(l1, l2)

	fmt.Println("\noutput:", Serialize(ans))
}
