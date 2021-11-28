package main

import "interfaces/tic-tac-toe/oop/cli"

func main() {
	cliGame := cli.NewGame([2]cli.Player{
		cli.NewHumanPlayer("X"), cli.NewAIPlayer(),
	})
	cliGame.Start()
}
