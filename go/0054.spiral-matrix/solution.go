// Created by lotus at 2025/03/07 09:21
// leetgo: 1.4.13
// https://leetcode.com/problems/spiral-matrix/

package main

import "fmt"

// import (
// 	"bufio"
// 	"fmt"
// 	"os"

// 	. "github.com/j178/leetgo/testutils/go"
// )

// @lc code=begin

func spiralOrder(matrix [][]int) (ans []int) {
	// top down left right ref the row that has not traverse
	top := 0
	down := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for top <= down && left <= right {
		// traverse top row
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++

		// traverse righ row
		for i := top; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		// traverse down low
		if top <= down {
			for i := right; i >= left; i-- {
				ans = append(ans, matrix[down][i])
			}
			down--
		}

		if left <= right {
			for i := down; i >= top; i-- {
				ans = append(ans, matrix[i][left])
			}
			left++
		}
	}
	return ans
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// matrix := Deserialize[[][]int](ReadLine(stdin))
	// ans := spiralOrder(matrix)

	// fmt.Println("\noutput:", Serialize(ans))

	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(spiralOrder(m))
}
