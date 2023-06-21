package game

import (
	"Assignment1/TicTacToeStruct/board"
	"Assignment1/TicTacToeStruct/player"
	"errors"
	// "fmt"
)

type Game struct {
	Players []*player.Player
	board   *board.Board
	turn    int
}

func NewGame(P1Name, P2Name string) *Game {
	player1 := player.NewPlayer(P1Name, "X")
	player2 := player.NewPlayer(P2Name, "O")
	boardCreated := board.NewBoard()

	return &Game{
		Players: []*player.Player{player1, player2},
		board:   boardCreated,
		turn:    0,
	}
}

func (g *Game) Play(cellNum int) (string, error) {
	if g.turn == -1 {
		return "", errors.New("Game has ended")
	}

	if g.board.Cells[cellNum].IsCellMarked() {
		return "", errors.New("cell already marked")
	}

	currentPlayer := g.Players[g.turn%2]
	g.board.Cells[cellNum].MarkCell(currentPlayer.Symbol)
	g.turn++

	g.board.ShowBoard()

	if g.board.ResultAnalyzer() == 1 {
		g.turn = -1
		//fmt.Println(currentPlayer.Name,"(",currentPlayer.Symbol,")","won the game")
		result := currentPlayer.Name + " (" + currentPlayer.Symbol + ") " + "won the game\nThe game has ended"
		return result, nil
	}

	return "", nil
}
