package main

import (
	"fmt"
)

var turn byte = 'o'
var computer byte = 'o'
var user byte = 'x'

func main() {
	var state = [3][3]byte{}
	fmt.Println(run(state))
}

func run(state [3][3]byte) int {
	if isWinningState(state, 'x') {
		fmt.Print("x wins ")
		fmt.Println(state)
		return 1
	} else if isWinningState(state, 'o') {
		fmt.Print("o wins ")
		fmt.Println(state)
		return -1
	} else {
		fmt.Println(state)
		if turn == computer {
			fs := getFutureStates(state, turn)
			switchTurn()
			for _, s := range fs {
				return run(s)
			}
		} else {
			state = getUserInput(state)
			switchTurn()
			run(state)
		}
	}

	return 0
}

func isWinningState(board [3][3]byte, t byte) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == t && board[i][1] == t && board[i][2] == t {
			return true
		}

		if board[0][i] == t && board[1][i] == t && board[2][i] == t {
			return true
		}
	}
	if (board[0][0] == t && board[1][1] == t && board[2][2] == t) ||
		(board[0][2] == t && board[1][1] == t && board[2][0] == t) {
		return true
	}

	return false
}

func getFutureStates(state [3][3]byte, turn byte) [9][3][3]byte {
	result := [9][3][3]byte{}
	k := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if state[i][j] == 0 {
				state[i][j] = turn
				result[k] = state
				k++
				break
			}
		}
	}

	return result
}

func switchTurn() {
	if turn == 'x' {
		turn = 'o'
	} else {
		turn = 'x'
	}
}

func getUserInput(state [3][3]byte) [3][3]byte {
	var i, j int
	fmt.Print("Enter row: ")
	fmt.Scanf("%d", &i)
	fmt.Print("Enter column: ")
	fmt.Scanf("%d", &j)
	state[i][j] = user
	return state
}
