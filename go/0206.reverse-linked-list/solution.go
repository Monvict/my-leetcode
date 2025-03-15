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

	var tail *ListNode
	current := head

	for current != nil {
		// 先记录下一个节点（因为下一步current.Next就改变了）
		next := current.Next
		current.Next = tail

		// tail节点后移
		tail = current
		// 当前节点后移
		current = next
	}

	return tail
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
