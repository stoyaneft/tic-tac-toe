package cli

import (
	"fmt"
	"interfaces/tic-tac-toe/oop/game"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type Player interface {
	Play(board []byte) int
	Name() string
}

type HumanPlayer struct {
	name string
}

func (p *HumanPlayer) Play(board []byte) int {
	fmt.Printf("What is your move player %s (1-9)?\n", p.Name())
	var move int
	fmt.Scan(&move)
	return move
}

func (p *HumanPlayer) Name() string {
	return p.name
}

func NewHumanPlayer(name string) Player {
	return &HumanPlayer{
		name: name,
	}
}

type AIPlayer struct{}

func (p *AIPlayer) Play(board []byte) int {
	return rand.Int() % len(board)
}

func (p *AIPlayer) Name() string {
	return "AI"
}

func NewAIPlayer() Player {
	rand.Seed(time.Now().Unix())
	return &AIPlayer{}
}

type CLI struct {
	players [2]Player
	turn    int
}

func NewGame(players [2]Player) *CLI {
	return &CLI{
		players: players,
		turn:    0,
	}
}

func (c *CLI) Start() {
	g := game.New()
	result := game.IN_PROGRESS

	for result == game.IN_PROGRESS {
		c.drawBoard(g.GetBoard())
		result = c.playMove(g)
	}

	c.drawBoard(g.GetBoard())

	switch result {
	case game.X_WINS:
		fmt.Println("Congrats X wins!")
	case game.O_WINS:
		fmt.Println("Congrats O wins!")
	case game.TIE:
		fmt.Println("It's a tie!")
	}
}

func (c *CLI) playMove(g *game.Game) game.Result {
	currentPlayer := c.players[c.turn%len(c.players)]
	board := g.GetBoard()
	move := currentPlayer.Play(board)
	result, err := g.Play(move)
	for err == game.ErrInvalidMove {
		fmt.Println("Illegal move! Play again.")
		move := currentPlayer.Play(board)
		result, err = g.Play(move)
	}
	c.turn++
	return result
}

func (c *CLI) drawBoard(board []byte) {
	fmt.Print("\033[H\033[2J") // clear screen (NOTE: not cross-platfrom)
	newBoard := re.ReplaceAllFunc(BOARD_TEMPLATE, func(s []byte) []byte {
		i, _ := strconv.Atoi(string(s))
		return []byte{board[i]}
	})
	fmt.Printf("%s\n", newBoard)
}

var re = regexp.MustCompile(`\d`)
var BOARD_TEMPLATE = []byte(`
 0 │ 1 │ 2
───┼───┼───
 3 │ 4 │ 5
───┼───┼───
 6 │ 7 │ 8
`)
