// Created by lotus at 2025/02/18 22:35
// leetgo: 1.4.13
// https://leetcode.com/problems/merge-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 1. 从后面开始往前填值
	tail := m + n - 1
	idx1 := m - 1
	idx2 := n - 1

	for {
		// all done
		if idx1 < 0 && idx2 < 0 {
			break
		}

		// nums2 left
		if idx2 >= 0 && idx1 < 0 {
			nums1[tail] = nums2[idx2]
			tail--
			idx2--
			continue
		}

		if idx1 >= 0 && idx2 < 0 {
			nums1[tail] = nums1[idx1]
			tail--
			idx1--
			continue
		}

		if idx1 >= 0 && idx2 >= 0 {
			if nums1[idx1] > nums2[idx2] {
				nums1[tail] = nums1[idx1]
				idx1--
			} else {
				nums1[tail] = nums2[idx2]
				idx2--
			}
			tail--
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	m := Deserialize[int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	merge(nums1, m, nums2, n)
	ans := nums1

	fmt.Println("\noutput:", Serialize(ans))
}
