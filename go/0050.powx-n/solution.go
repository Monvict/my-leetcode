// Created by lotus at 2025/03/13 12:33
// leetgo: 1.4.13
// https://leetcode.com/problems/powx-n/

package main

import (
	"fmt"
	"math"
)

// @lc code=begin
func myPow(x float64, n int) (ans float64) {

	result := recursion(x, int(math.Abs(float64(n))))

	if n < 0 {
		return 1 / result
	}

	return result
}

// dived and quanqer so the O(n) = log(n)
func recursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	half := recursion(x, n>>1)
	if n&1 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

// @lc code=end

func main() {
	fmt.Println(myPow(2, 5))
}
