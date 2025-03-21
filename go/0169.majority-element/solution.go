// Created by lotus at 2025/03/20 23:24
// leetgo: 1.4.13
// https://leetcode.com/problems/majority-element/

package main

import "fmt"

// @lc code=begin

func majorityElement(nums []int) (ans int) {

	// refer https://www.cs.utexas.edu/~moore/best-ideas/mjrty/index.html
	cnt := 1
	majority := nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		if cnt == 0 {
			majority = num
		}

		if num == majority {
			cnt++
		} else {
			cnt--
		}

	}
	return majority
}

func majorityElement2(nums []int) (ans int) {
	numMap := make(map[int]int)
	len := len(nums)
	for _, num := range nums {
		numMap[num]++

		if numMap[num] > len/2 {
			return num
		}
	}
	return 0
}

// @lc code=end

func main() {
	number := []int{6, 5, 5}

	fmt.Println("Maj ", majorityElement(number))
}
