// Created by lotus at 2025/03/07 08:51
// leetgo: 1.4.13
// https://leetcode.com/problems/matrix-diagonal-sum/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func diagonalSum(mat [][]int) (ans int) {

	sum := 0
	for i, m := range mat {
		firstEle := m[i]
		secondEle := m[len(m)-i-1]
		sum += firstEle
		sum += secondEle
	}

	if len(mat)%2 != 0 {
		mid := len(mat) / 2
		midEle := mat[mid][mid]
		sum -= midEle
	}

	return sum
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	mat := Deserialize[[][]int](ReadLine(stdin))
	ans := diagonalSum(mat)

	fmt.Println("\noutput:", Serialize(ans))

	// mat := [][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }

	// fmt.Println(diagonalSum(mat))
}
