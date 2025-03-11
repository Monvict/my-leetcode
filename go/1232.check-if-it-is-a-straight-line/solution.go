// Created by lotus at 2025/03/11 22:32
// leetgo: 1.4.13
// https://leetcode.com/problems/check-if-it-is-a-straight-line/

package main

import "fmt"

// @lc code=begin

func checkStraightLine(coordinates [][]int) bool {
	theta := calcTheta(coordinates[0], coordinates[1])

	for i := 1; i < len(coordinates)-1; i++ {
		t := calcTheta(coordinates[i], coordinates[i+1])

		if theta != t {
			return false
		}
	}

	return true
}

func calcTheta(a []int, b []int) float64 {
	deltaY := b[1] - a[1]
	deltaX := b[0] - a[0]

	if deltaX == 0 {
		return float64(999999)
	}

	return float64(deltaY) / float64(deltaX)
}

// @lc code=end

func main() {
	c := [][]int{
		{-2, 2},
		{-1, 1},
		{0, 0},
		{0, 1},
		{0, -1},
	}

	fmt.Println(checkStraightLine(c))
}
