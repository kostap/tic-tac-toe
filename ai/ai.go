// Simple Tic-Tac-Toe AI

package ai

import (
	"github.com/kostap/tic-tac-toe/game"
	"math/rand/v2"
)

type AI struct {
	Player game.Player
}

func NewAI(player game.Player) *AI {
	return &AI{
		Player: player,
	}
}

func (ai *AI) MakeMove(g *game.Game) (int, int) {
	// Try to find winner position
	if row, col := ai.findWinningMove(g, ai.Player); row != -1 {
		return row, col
	}

	// Block opponent move
	opponent := game.PlayerX
	if ai.Player == game.PlayerX {
		opponent = game.PlayerO
	}
	if row, col := ai.findWinningMove(g, opponent); row != -1 {
		return row, col
	}

	// Occupy the center if empty
	if g.Board[1][1] == game.Empty {
		return 1, 1
	}

	// Occupy the corners
	corners := [][2]int{{0, 0}, {0, 2}, {2, 0}, {2, 2}}
	rand.Shuffle(len(corners), func(i, j int) {
		corners[i], corners[j] = corners[j], corners[i]
	})
	for _, corner := range corners {
		if g.Board[corner[0]][corner[1]] == game.Empty {
			return corner[0], corner[1]
		}
	}

	// Random Move
	return ai.randomMove(g)
}

func (ai *AI) findWinningMove(g *game.Game, player game.Player) (int, int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == game.Empty {
				g.Board[i][j] = player

				if g.CheckWinner() {
					g.Board[i][j] = game.Empty

					return i, j
				}

				g.Board[i][j] = game.Empty
			}
		}
	}

	return -1, -1 // Invalid positions
}

func (ai *AI) randomMove(g *game.Game) (int, int) {
	available := make([][2]int, 0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == game.Empty {
				available = append(available, [2]int{i, j})
			}
		}
	}

	if len(available) == 0 {
		return -1, -1 // Invalid positions
	}

	move := available[rand.IntN(len(available))]

	return move[0], move[1]
}
