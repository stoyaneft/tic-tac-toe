package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	playGame()
}

func playGame() {
	startBoard := []byte{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY}
	players := []byte{X, O}

	board := startBoard
	result := IN_PROGRESS
	turn := 0

	for result == IN_PROGRESS {
		currentPlayer := players[turn%len(players)]
		drawBoard(board)
		err := playMove(currentPlayer, board)
		for err == ErrInvalidMove {
			fmt.Println("Illegal move! Play again.")
			err = playMove(currentPlayer, board)
		}
		result = getResult(board)
		turn++
	}

	drawBoard(board)

	switch result {
	case X_WINS:
		fmt.Println("Congrats X wins!")
	case O_WINS:
		fmt.Println("Congrats O wins!")
	case TIE:
		fmt.Println("It's a tie!")
	}

}

func getResult(board []byte) GameResult {
	var winner *byte

	for _, line := range WINNING_LINES {
		player := board[line[0]]

		isWinning := true
		for _, i := range line {
			if player != board[i] {
				isWinning = false
				break
			}
		}
		if isWinning {
			winner = &player
			break
		}
	}

	if winner == nil {
		if areEmptyCells(board) {
			return IN_PROGRESS
		}
		return TIE
	}

	switch *winner {
	case X:
		return X_WINS
	case O:
		return O_WINS
	}
	return IN_PROGRESS
}

func areEmptyCells(board []byte) bool {
	for _, b := range board {
		if b == EMPTY {
			return true
		}
	}
	return false
}

func playMove(player byte, board []byte) error {
	fmt.Printf("What is your move player %q (1-9)?\n", player)
	var position int
	fmt.Scan(&position)
	position--

	if !isValidMove(position, board) {
		return ErrInvalidMove
	}
	board[position] = player
	return nil
}

func isValidMove(position int, board []byte) bool {
	return 0 <= position && position < len(board) && board[position] == EMPTY
}

func drawBoard(board []byte) {
	fmt.Print("\033[H\033[2J") // clear screen (NOTE: not cross-platfrom)
	newBoard := re.ReplaceAllFunc(BOARD_TEMPLATE, func(s []byte) []byte {
		i, _ := strconv.Atoi(string(s))
		return []byte{board[i]}
	})
	fmt.Printf("%s\n", newBoard)
}

var BOARD_TEMPLATE = []byte(`
 0 │ 1 │ 2
───┼───┼───
 3 │ 4 │ 5
───┼───┼───
 6 │ 7 │ 8
`)
var WINNING_LINES = [8][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},

	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},

	{0, 4, 8},
	{6, 4, 2},
}
var re = regexp.MustCompile(`\d`)

var ErrInvalidMove = fmt.Errorf("invalid move")

type GameResult int

const (
	IN_PROGRESS GameResult = iota
	X_WINS
	O_WINS
	TIE
)

const (
	X     = 'x'
	O     = 'o'
	EMPTY = ' '
)
