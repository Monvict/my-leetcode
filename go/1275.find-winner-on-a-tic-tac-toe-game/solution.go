// Created by lotus at 2025/02/28 14:48
// leetgo: 1.4.13
// https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func tictactoe(moves [][]int) string {
	// init a matrix
	board := make([][]int, 3)
	for i := range board {
		board[i] = make([]int, 3)
	}

	// put the moves in a game board
	for i, mov := range moves {
		row := mov[0]
		col := mov[1]

		// 1 present for A
		// 2 present for B
		if i%2 == 0 {
			board[row][col] = 1
		} else {
			board[row][col] = 2
		}

		// check for has a winner after a step
		winner := hasWinner(board)
		if winner == 1 {
			return "A"
		}

		if winner == 2 {
			return "B"
		}
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

func hasWinner(board [][]int) int {
	// row winner
	for i := 0; i < 3; i++ {
		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
	}

	// col winner
	for j := 0; j < 3; j++ {
		if board[0][j] != 0 && board[0][j] == board[1][j] && board[1][j] == board[2][j] {
			return board[0][j]
		}
	}

	// dia winner
	if board[0][0] != 0 && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}

	// another-dia winner
	if board[2][0] != 0 && board[2][0] == board[1][1] && board[1][1] == board[0][2] {
		return board[2][0]
	}

	return 0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	moves := Deserialize[[][]int](ReadLine(stdin))
	ans := tictactoe(moves)

	fmt.Println("\noutput:", Serialize(ans))

	// m := [][]int{
	// 	{0, 0},
	// 	{1, 1},
	// 	{2, 0},
	// 	{1, 0},
	// 	{1, 2},
	// 	{2, 1},
	// 	{0, 1},
	// 	{0, 2},
	// 	{2, 2},
	// }

	// fmt.Println(tictactoe(m))
}
