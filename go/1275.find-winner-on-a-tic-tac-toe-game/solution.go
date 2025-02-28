// Created by lotus at 2025/02/28 14:48
// leetgo: 1.4.13
// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/

package main

import "fmt"

// import (
// 	"bufio"
// 	"fmt"
// 	"os"

// 	. "github.com/j178/leetgo/testutils/go"
// )

// @lc code=begin

func tictactoe(moves [][]int) string {
	playerA_X := 0
	playerA_Y := 0
	playerB_X := 0
	playerB_Y := 0

	for i, mov := range moves {
		if i%2 == 0 {
			playerA_X += mov[0]
			playerA_Y += mov[1]
		} else {
			playerB_X += mov[0]
			playerB_Y += mov[1]
		}

		if playerA_X == 3 && playerA_Y == 3 {
			return "A"
		}

		if playerB_X == 3 && playerB_Y == 3 {
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

// @lc code=end

func main() {
	// stdin := bufio.NewReader(os.Stdin)
	// moves := Deserialize[[][]int](ReadLine(stdin))
	// ans := tictactoe(moves)

	// fmt.Println("\noutput:", Serialize(ans))

	m := [][]int{
		{0, 0},
		{1, 1},
		{2, 0},
		{1, 0},
		{1, 2},
		{2, 1},
		{0, 1},
		{0, 2},
		{2, 2},
	}

	fmt.Println(tictactoe(m))
}
