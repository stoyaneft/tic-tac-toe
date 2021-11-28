package game

import "fmt"

type Game struct {
	board         []byte
	turn          int
	currentPlayer byte
}

func New() *Game {
	return &Game{
		board:         []byte{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		currentPlayer: X,
		turn:          0,
	}
}

func (g *Game) Play(position int) (Result, error) {
	position--

	if !g.isValidMove(position) {
		return IN_PROGRESS, ErrInvalidMove
	}
	g.board[position] = g.currentPlayer
	g.turn++
	g.currentPlayer = players[g.turn%len(players)]
	return g.getResult(), nil
}

func (g *Game) GetBoard() []byte {
	var board = make([]byte, len(g.board))
	copy(board, g.board)
	return board
}

func (g *Game) isValidMove(position int) bool {
	return 0 <= position && position < len(g.board) && g.board[position] == EMPTY
}

func (g *Game) getResult() Result {
	var winner *byte

	for _, line := range WINNING_LINES {
		player := g.board[line[0]]

		isWinning := true
		for _, i := range line {
			if player != g.board[i] {
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
		if g.emptyCellsExist() {
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

func (g *Game) emptyCellsExist() bool {
	for _, b := range g.board {
		if b == EMPTY {
			return true
		}
	}
	return false
}

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

var players = [2]byte{X, O}

var ErrInvalidMove = fmt.Errorf("invalid move")

type Result int

const (
	IN_PROGRESS Result = iota
	X_WINS
	O_WINS
	TIE
)

const (
	X     = 'x'
	O     = 'o'
	EMPTY = ' '
)
