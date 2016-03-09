package main

import (
	"fmt"
	"os"
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
		state.printState()
		if state.isWinningState('o') {
			fmt.Println("Computer Wins!!")
			break
		}

		state.checkDraw()

		state.getUserInput()
		state.printState()
		if state.isWinningState('x') {
			fmt.Println("You Win!!")
			break
		}

		state.checkDraw()
	}
}

func run(state State, level int, turn byte, ch chan State) int {
	if state.isWinningState(computer) {
		return 1
	}
	if state.isWinningState(user) {
		return -1
	}

	min := 1
	max := -1
	var nextState State

	fs := state.getFutureStates(turn)

	for _, s := range fs {
		w := run(s, level+1, switchTurn(turn), nil)

		if turn == computer && w > max {
			max = w
			nextState = s
		} else if turn == user && w < min {
			min = w
			nextState = s
		}
	}

	if level == 0 {
		ch <- nextState
	}

	if turn == computer {
		return max
	} else {
		return min
	}
}

func (state State) isWinningState(t byte) bool {
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

func (state State) checkDraw() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if state[i][j] == 0 {
				return
			}
		}
	}

	os.Exit(0)
}

func (state State) getFutureStates(turn byte) []State {
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

	fmt.Println("Match Draw")
	return result
}

func switchTurn(turn byte) byte {
	if turn == 'x' {
		return 'o'
	} else {
		return 'x'
	}
}

func (state *State) getUserInput() {
	var i, j int

	for {
		fmt.Print("Enter row (0-2): ")
		fmt.Scanf("%d", &i)
		fmt.Print("Enter column (0-2): ")
		fmt.Scanf("%d", &j)
		if state[i][j] == 0 {
			state[i][j] = user
			return
		} else {
			fmt.Println("Invalid Move")
		}
	}
}

func (s State) printState() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%c\t", s[i][j])
		}
		fmt.Println()
	}
	fmt.Println("==================")
}
