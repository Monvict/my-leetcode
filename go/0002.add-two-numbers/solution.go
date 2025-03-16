// Created by lotus at 2025/03/16 22:26
// leetgo: 1.4.13
// https://leetcode.com/problems/add-two-numbers/

package main

import (
	"fmt"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ans *ListNode) {

	a := 0
	b := 0
	carry := 0

	r := &ListNode{}
	pr := r
	for l1 != nil && l2 != nil {
		a = l1.Val
		b = l2.Val

		sum := a + b + carry
		cur := sum % 10
		carry = sum / 10

		node := &ListNode{
			Val: cur,
		}
		pr.Next = node
		pr = pr.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	if l1 != nil {
		if carry > 0 {
			addCarry(l1)
		}

		pr.Next = l1
	}

	if l2 != nil {
		if carry > 0 {
			addCarry(l2)
		}

		pr.Next = l2
	}

	if l1 == nil && l2 == nil && carry > 0 {
		pr.Next = &ListNode{Val: 1}
	}
	return r.Next
}

func addCarry(l *ListNode) {
	p := l
	carry := 1
	for l != nil {
		val := l.Val
		sum := val + carry

		cur := sum % 10
		carry = sum / 10

		l.Val = cur
		l = l.Next
	}

	if carry > 0 {
		for p.Next != nil {
			p = p.Next
		}

		p.Next = &ListNode{
			Val: 1,
		}
	}
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
	p := createList([]int{5})
	p2 := createList([]int{5})

	r := addTwoNumbers(p, p2)
	fmt.Println(r.Values())
}
