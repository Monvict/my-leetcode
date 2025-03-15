// Created by lotus at 2025/03/15 09:17
// leetgo: 1.4.13
// https://leetcode.com/problems/reverse-linked-list/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func reverseList(head *ListNode) (ans *ListNode) {

	vals := make([]int, 0)
	for head != nil {
		val := head.Val
		vals = append(vals, val)
		head = head.Next
	}

	dummy := &ListNode{}
	p := dummy
	for i := len(vals) - 1; i >= 0; i-- {
		n := &ListNode{
			Val: vals[i],
		}
		p.Next = n
		p = p.Next
	}

	return dummy.Next
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{}
	p := head

	for _, num := range nums {
		node := &ListNode{
			Val:  num,
			Next: nil,
		}

		p.Next = node
		p = p.Next
	}

	return head.Next
}

// @lc code=end

func main() {
	l := createList([]int{3, 8, 2, 1})
	r := reverseList(l)
	fmt.Println(r)
}
