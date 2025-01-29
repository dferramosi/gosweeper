package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	rows, cols, mines := 8, 16, 20
	gameBoard := NewBoard(rows, cols, mines)

	assert.NotNil(t, gameBoard)
	assert.Equal(t, rows, gameBoard.Rows)
	assert.Equal(t, cols, gameBoard.Columns)
	assert.Equal(t, mines, gameBoard.Mines)
	assert.Equal(t, rows, len(gameBoard.Cells))
	assert.Equal(t, cols, len(gameBoard.Cells[0]))
}

func TestAddMine(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0)
	success, err := gameBoard.AddMine(0, 0)

	assert.NoError(t, err)
	assert.True(t, success)
	assert.True(t, gameBoard.Cells[0][0].IsMine)
	assert.Equal(t, 1, gameBoard.Mines)
}

func TestPopulateMines(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0)
	err := gameBoard.populateMines(20)

	assert.NoError(t, err)
	assert.Equal(t, 20, gameBoard.Mines)
}

func TestPopulateAdjacents(t *testing.T) {
	gameBoard := NewBoard(8, 16, 20)
	gameBoard.PopulateAdjacents()

	for row := 0; row < gameBoard.Rows; row++ {
		for col := 0; col < gameBoard.Columns; col++ {
			cell := gameBoard.Cells[row][col]
			if !cell.IsMine {
				adjacents, _ := gameBoard.getAdjacents(row, col)
				assert.Equal(t, adjacents, cell.Adj)
			}
		}
	}
}

func TestFlag(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0)
	gameBoard.Flag(0, 0)

	assert.True(t, gameBoard.Cells[0][0].IsFlag)
	assert.Equal(t, 1, gameBoard.Flags)

	gameBoard.Flag(0, 0)
	assert.False(t, gameBoard.Cells[0][0].IsFlag)
	assert.Equal(t, 0, gameBoard.Flags)
}

func TestReveal(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0)
	gameBoard.Cells[0][0].IsMine = false
	gameBoard.Cells[0][0].Adj = 1

	s, err := gameBoard.Reveal(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, "1", s)
	assert.True(t, gameBoard.Cells[0][0].IsRevealed)
	assert.Equal(t, 1, gameBoard.Revealed)
}

func TestRevealAdjacent(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0)
	gameBoard.Cells[0][0].Adj = 0
	gameBoard.Cells[0][1].Adj = 1
	gameBoard.Cells[1][0].Adj = 1
	gameBoard.Cells[1][1].Adj = 1

	visited := gameBoard.RevealAdjacent(0, 0, []string{}, true)
	assert.Contains(t, visited, "0,0")
	assert.Contains(t, visited, "0,1")
	assert.Contains(t, visited, "1,0")
	assert.Contains(t, visited, "1,1")
}
