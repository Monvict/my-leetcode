// Created by lotus at 2025/03/21 21:53
// leetgo: 1.4.13
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"fmt"
)

// @lc code=begin

func maxProfit(prices []int) (ans int) {
	// buyPrice := prices[0]
	// bueIndex := 0
	// for i := 1; i < len(prices); i++ {
	// 	if prices[i] < buyPrice {
	// 		buyPrice = prices[i]
	// 		bueIndex = i
	// 	}
	// }

	// salePrice := buyPrice
	// for j := bueIndex; j < len(prices); j++ {
	// 	if prices[j] > salePrice {
	// 		salePrice = prices[j]
	// 	}
	// }

	// profit := 0

	// for i := 0; i < len(prices); i++ {
	// 	for j := i + 1; j < len(prices); j++ {
	// 		delta := prices[j] - prices[i]
	// 		if delta > profit {
	// 			profit = delta
	// 		}
	// 	}
	// }

	profit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		delta := prices[i] - minPrice

		if delta < 0 {
			minPrice = prices[i]
		} else if delta > profit {
			profit = delta
		}
	}
	return profit
}

// @lc code=end

func main() {
	nums := []int{2, 4, 1}
	fmt.Println(maxProfit(nums))
}
