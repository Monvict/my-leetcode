// Created by lotus at 2025/03/26 22:33
// leetgo: 1.4.13
// https://leetcode.com/problems/linked-list-cycle/

package main

import (
	"fmt"
	"leetcode-solutions/mutils"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hasCycle(head *ListNode) bool {
	low := head
	fast := head

	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next

		if fast == low {
			return true
		}
	}

	return false
}

// @lc code=end

// Warning: this is a manual question, the generated test code may be incorrect.
func main() {
	fmt.Println(hasCycle(mutils.CreateList([]int{3, 2, 0, -4, 2})))
}
