// Created by lotus at 2025/03/10 20:28
// leetgo: 1.4.13
// https://leetcode.com/problems/lemonade-change/

package main

import "fmt"

// @lc code=begin
// 贪心算法：找钱时，尽量先找10元，保留5元，以备下此可以使用
// 当收到20元时：
// 1. 如果有10元，和5元的；找
// 2. 有3张以上5元，找开
// 3. 否则，找不开
func lemonadeChange(bills []int) bool {
	fiveCnt := 0
	tenCnt := 0

	for _, bill := range bills {
		if bill == 5 {
			fiveCnt++
		}

		if bill == 10 {
			// 找不开
			if fiveCnt <= 0 {
				return false
			}
			tenCnt++
			fiveCnt--
		}

		if bill == 20 {
			// 有10元和5元
			if tenCnt > 0 && fiveCnt > 0 {
				tenCnt--
				fiveCnt--
				// 有3个以上5元
			} else if fiveCnt > 2 {
				fiveCnt -= 3
			} else {
				// 找不开
				return false
			}
		}
	}

	return true
}

// @lc code=end

func main() {
	fmt.Println(lemonadeChange([]int{10, 10}))
}
