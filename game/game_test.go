package game

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame()

	if g.CurrentPlayer != PlayerX {
		t.Error("First player should be X")
	}
}

func TestMakeMove(t *testing.T) {
	g := NewGame()

	err := g.MakeMove(0, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if g.Board[0][0] != PlayerX {
		t.Error("Board position should be X")
	}
}

func TestWinCondition(t *testing.T) {
	g := NewGame()

	// X wins on first row
	g.MakeMove(0, 0) // X
	g.MakeMove(1, 0) // O
	g.MakeMove(0, 1) // X
	g.MakeMove(1, 1) // O
	g.MakeMove(0, 2) // X

	if !g.IsGameOver || g.Winner != PlayerX {
		t.Error("X should have won")
	}
}
