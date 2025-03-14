// Created by lotus at 2025/03/14 23:02
// leetgo: 1.4.13
// https://leetcode.com/problems/merge-two-sorted-lists/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	p := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}

		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return head.Next
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
	p := createList([]int{1, 3, 4})
	p2 := createList([]int{1, 2, 4})

	r := mergeTwoLists(p, p2)
	fmt.Println(r)
}
