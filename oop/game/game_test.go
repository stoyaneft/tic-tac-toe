package game_test

import (
	"interfaces/tic-tac-toe/oop/game"
	"testing"
)

func TestNewGame(t *testing.T) {
	g := game.New()

	board := g.GetBoard()
	for _, cell := range board {
		if cell != game.EMPTY {
			t.Errorf("starting board contains non-empty cell")
			return
		}
	}
}

func TestPlayGame(t *testing.T) {
	g := game.New()

	res, err := g.Play(1)
	assertResult(t, game.IN_PROGRESS, res)
	assertError(t, nil, err)

	res, err = g.Play(4)
	assertResult(t, game.IN_PROGRESS, res)
	assertError(t, nil, err)

	res, err = g.Play(2)
	assertResult(t, game.IN_PROGRESS, res)
	assertError(t, nil, err)

	res, err = g.Play(5)
	assertResult(t, game.IN_PROGRESS, res)
	assertError(t, nil, err)

	res, err = g.Play(5)
	assertResult(t, game.IN_PROGRESS, res)
	assertError(t, game.ErrInvalidMove, err)

	res, err = g.Play(3)
	assertResult(t, game.X_WINS, res)
	assertError(t, nil, err)

}

func assertResult(t *testing.T, expectedResult game.Result, actualResult game.Result) {
	if actualResult != expectedResult {
		t.Errorf("expected result to be %d but got %d", expectedResult, actualResult)
	}
}

func assertError(t *testing.T, expectedError error, actualError error) {
	if actualError != expectedError {
		t.Errorf("expected error to be %s but got %s", expectedError, actualError)
	}
}
