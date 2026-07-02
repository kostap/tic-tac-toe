package handler

import (
	"encoding/json"
	"github.com/kostap/tic-tac-toe/internal/ai"
	"github.com/kostap/tic-tac-toe/internal/game"
	"net/http"
)

type GameHandler struct {
	currentGame *game.Game
	ai          *ai.AI
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		currentGame: game.NewGame(),
		ai:          ai.NewAI(game.PlayerO),
	}
}

type MoveRequest struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

type GameResponse struct {
	Board         [][]string `json:"board"`
	CurrentPlayer string     `json:"currentPlayer"`
	Winner        string     `json:"winner"`
	IsDraw        bool       `json:"isDraw"`
	GameOver      bool       `json:"gameOver"`
	Message       string     `json:"message"`
}

func (h *GameHandler) NewGame(w http.ResponseWriter, r *http.Request) {
	h.currentGame = game.NewGame()
	h.sendGameState(w)
}

func (h *GameHandler) GetState(w http.ResponseWriter, r *http.Request) {
	h.sendGameState(w)
}

func (h *GameHandler) MakeMove(w http.ResponseWriter, r *http.Request) {
	var moveReq MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&moveReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.currentGame.MakeMove(moveReq.Row, moveReq.Col); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !h.currentGame.IsGameOver {
		row, col := h.ai.MakeMove(h.currentGame)
		if row != -1 && col != -1 {
			_ = h.currentGame.MakeMove(row, col)
		}
	}

	h.sendGameState(w)
}

func (h *GameHandler) sendGameState(w http.ResponseWriter) {
	response := GameResponse{
		CurrentPlayer: string(h.currentGame.CurrentPlayer),
		Winner:        string(h.currentGame.Winner),
		IsDraw:        h.currentGame.IsDraw,
		GameOver:      h.currentGame.IsGameOver,
		Board:         make([][]string, 3),
	}

	for i := 0; i < 3; i++ {
		response.Board[i] = make([]string, 3)
		for j := 0; j < 3; j++ {
			response.Board[i][j] = string(h.currentGame.Board[i][j])
		}
	}

	if response.GameOver {
		if response.IsDraw {
			response.Message = "No one won!"
		} else {
			response.Message = "Winner: " + response.Winner
		}
	} else {
		response.Message = "Player's turn: " + response.CurrentPlayer
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}
