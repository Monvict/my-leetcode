// Created by lotus at 2025/03/09 22:14
// leetgo: 1.4.13
// https://leetcode.com/problems/set-matrix-zeroes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func setZeroes(matrix [][]int) {
	// 空间复杂度O(1)的解题思路
	// 1. 第一行，第一列用来记录某行、某列是否需要置零
	// 2. 因为第一行、第一列用来记录，所以第一行，第一列自己须否需要置零，则需要在最开始用2个标志位记录
	// 3. 遍历除第一行、第一列之外的的元素，当碰到0时，则在第一行、第一列中记录0
	// 4. 第一行、第一列记录完成后，再根据零的值将响应的行列置零
	// 5. 最后，根据标志位将第一行、或第一列置零

	firstRowZero := false
	firstColumnnZero := false

	// 标志位记录第一行、第一列是否要置零
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			firstColumnnZero = true
			break
		}
	}

	// 遍历除第一行、第一列之外的元素
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			// 碰到0，则在第一行中进行设置
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	// 根据第一行的标志，将0对应的列置零
	for j := 1; j < len(matrix[0]); j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}

	// 根据第一列的值，将0对应的行置零
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < len(matrix[0]); j++ {
				matrix[i][j] = 0
			}
		}
	}

	// 第一行置零
	if firstRowZero {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}

	// 第一列置零
	if firstColumnnZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	setZeroes(matrix)
	ans := matrix

	fmt.Println("\noutput:", Serialize(ans))
}
