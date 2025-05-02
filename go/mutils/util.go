package mutils

import . "github.com/j178/leetgo/testutils/go"

func CreateList(nums []int) *ListNode {
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
