package main

import (
	"fmt"
)

type State [3][3]byte

var computer byte = 'o'
var user byte = 'x'

func main() {
	var state = State{}

	for {
		ch := make(chan State, 9)
		run(state, 0, computer, ch)
		state = <-ch
		printState(state)
		if isWinningState(state, 'o') {
			fmt.Println("Computer Wins!!")
			break
		}

		state = getUserInput(state)
		printState(state)
		if isWinningState(state, 'x') {
			fmt.Println("You Win!!")
			break
		}
	}
}

func run(state State, level int, turn byte, ch chan State) int {
	if isWinningState(state, 'x') {
		return -1
	} else if isWinningState(state, 'o') {
		return 1
	} else {
		fs := getFutureStates(state, turn)
		for _, s := range fs {
			w := run(s, level+1, switchTurn(turn), ch)
			if w == 1 {
				//printState(s)
				if level == 0 {
					ch <- s
				} else {
					return 1
				}
			} else {
				return w
			}
		}
	}
	return 0
}

func isWinningState(state State, t byte) bool {
	for i := 0; i < 3; i++ {
		if state[i][0] == t && state[i][1] == t && state[i][2] == t {
			return true
		}

		if state[0][i] == t && state[1][i] == t && state[2][i] == t {
			return true
		}
	}
	if (state[0][0] == t && state[1][1] == t && state[2][2] == t) ||
		(state[0][2] == t && state[1][1] == t && state[2][0] == t) {
		return true
	}

	return false
}

func getFutureStates(state State, turn byte) []State {
	result := []State{}
	s := state

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if s[i][j] == 0 {
				s[i][j] = turn
				ts := state
				ts[i][j] = turn
				result = append(result, ts)
			}
		}
	}

	return result
}

func switchTurn(turn byte) byte {
	if turn == 'x' {
		return 'o'
	} else {
		return 'x'
	}
}

func getUserInput(state State) State {
	var i, j int
	fmt.Print("Enter row: ")
	fmt.Scanf("%d", &i)
	fmt.Print("Enter column: ")
	fmt.Scanf("%d", &j)
	state[i][j] = user
	return state
}

func printState(state State) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c\t", state[i][j])
		}
		fmt.Println()
	}
	fmt.Println("==================")
}
