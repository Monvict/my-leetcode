// Created by lotus at 2025/03/06 21:23
// leetgo: 1.4.13
// https://leetcode.com/problems/robot-bounded-in-circle/

package main

import "fmt"

// import (
// 	"bufio"
// 	"fmt"
// 	"os"

// 	. "github.com/j178/leetgo/testutils/go"
// )

// @lc code=begin

func isRobotBounded(instructions string) bool {

	// 4个方向
	directions := [][]int{
		{0, 1},  // N(0)
		{-1, 0}, // W(1)
		{0, -1}, // S(2)
		{1, 0},  // E(3)
	}

	// 初始方向：N(0)
	dir := 0
	x, y := 0, 0

	for _, ins := range instructions {

		// 计算一下走一步移动的距离
		if ins == 'G' {
			dx := directions[dir][0]
			dy := directions[dir][1]

			x += dx
			y += dy
		} else if ins == 'L' {
			// 左转，计算dir的值
			dir = (dir + 1) % 4
		} else {
			dir = (dir + 3) % 4
		}
	}

	// 会形成环有2个条件
	// 1. 回到原点
	// 2. 不回到原点，并且方向不是被
	return (x == 0 && y == 0) || dir != 0
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// instructions := Deserialize[string](ReadLine(stdin))
	// ans := isRobotBounded(instructions)

	// fmt.Println("\noutput:", Serialize(ans))

	fmt.Println(isRobotBounded("GLRLGLLGLGRGLGL"))
}
