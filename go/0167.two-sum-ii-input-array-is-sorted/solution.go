// Created by lotus at 2025/03/29 23:40
// leetgo: 1.4.13
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

package main

import "fmt"

// @lc code=begin

func twoSum(numbers []int, target int) (ans []int) {
	low, high := 0, len(numbers)-1
	r := make([]int, 0)
	for low < high {
		low_val := numbers[low]
		high_val := numbers[high]

		if low_val+high_val == target {
			r = append(r, low+1, high+1)
			return r
		} else if low_val+high_val > target {
			high--
		} else {
			low++
		}
	}
	return r
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// numbers := Deserialize[[]int](ReadLine(stdin))
	// target := Deserialize[int](ReadLine(stdin))
	// ans := twoSum(numbers, target)

	// fmt.Println("\noutput:", Serialize(ans))

	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
