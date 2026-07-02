package game

import (
	"fmt"
)

type Player string

const (
	PlayerX Player = "X"
	PlayerO Player = "O"
	Empty   Player = ""
)

type Board [3][3]Player

type Game struct {
	Board         Board
	CurrentPlayer Player
	Winner        Player
	IsDraw        bool
	IsGameOver    bool
}

func NewGame() *Game {
	return &Game{
		CurrentPlayer: PlayerX,
		Board:         Board{},
	}
}

func (g *Game) MakeMove(row, col int) error {
	// Check if game ended
	if g.IsGameOver {
		return fmt.Errorf("game is over")
	}

	// Validate position
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return fmt.Errorf("invalid position")
	}

	// Check position
	if g.Board[row][col] != Empty {
		return fmt.Errorf("position already taken")
	}

	g.Board[row][col] = g.CurrentPlayer

	// Check Winner
	if g.CheckWinner() {
		g.Winner = g.CurrentPlayer
		g.IsGameOver = true

		return nil
	}

	// Check that we can make the next move in the game
	if g.CheckDraw() {
		g.IsDraw = true
		g.IsGameOver = true

		return nil
	}

	// Change current player
	g.ChangePlayer()

	return nil
}

func (g *Game) ChangePlayer() {
	if g.CurrentPlayer == PlayerX {
		g.CurrentPlayer = PlayerO
	} else {
		g.CurrentPlayer = PlayerX
	}
}

func (g *Game) CheckWinner() bool {
	// Check rows and cols
	for i := 0; i < 3; i++ {
		if g.Board[i][0] != Empty && g.Board[i][0] == g.Board[i][1] && g.Board[i][1] == g.Board[i][2] {
			return true
		}
		if g.Board[0][i] != Empty && g.Board[0][i] == g.Board[1][i] && g.Board[1][i] == g.Board[2][i] {
			return true
		}
	}

	// Check the diagonals
	if g.Board[0][0] != Empty && g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] {
		return true
	}
	if g.Board[0][2] != Empty && g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0] {
		return true
	}

	return false
}

func (g *Game) CheckDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == Empty {
				return false
			}
		}
	}

	return true
}
