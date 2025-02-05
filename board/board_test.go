package board

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func mockClickHandler() bool {
	return true
}

func TestNewBoard(t *testing.T) {
	rows, cols, mines := 8, 16, 20
	gameBoard := NewBoard(rows, cols, mines, mockClickHandler)

	assert.NotNil(t, gameBoard)
	assert.Equal(t, rows, gameBoard.Rows)
	assert.Equal(t, cols, gameBoard.Columns)
	assert.Equal(t, mines, gameBoard.Mines)
	assert.Equal(t, rows, len(gameBoard.Cells))
	assert.Equal(t, cols, len(gameBoard.Cells[0]))
}

func TestAddMine(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0, mockClickHandler)
	success, err := gameBoard.AddMine(0, 0)

	assert.NoError(t, err)
	assert.True(t, success)
	assert.True(t, gameBoard.Cells[0][0].IsMine)
	assert.Equal(t, 1, gameBoard.Mines)
}

func TestPopulateMines(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0, mockClickHandler)
	err := gameBoard.populateMines(20)

	assert.NoError(t, err)
	assert.Equal(t, 20, gameBoard.Mines)
}

func TestPopulateAdjacents(t *testing.T) {
	gameBoard := NewBoard(8, 16, 20, mockClickHandler)
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
	gameBoard := NewBoard(8, 16, 0, mockClickHandler)
	gameBoard.Flag(0, 0)

	assert.True(t, gameBoard.Cells[0][0].IsFlag)
	assert.Equal(t, 1, gameBoard.Flags)

	gameBoard.Flag(0, 0)
	assert.False(t, gameBoard.Cells[0][0].IsFlag)
	assert.Equal(t, 0, gameBoard.Flags)
}

func TestReveal(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0, mockClickHandler)
	gameBoard.Cells[0][0].IsMine = false
	gameBoard.Cells[0][0].Adj = 1
	gameBoard.Cells[0][1].IsFlag = true
	gameBoard.Cells[0][2].IsRevealed = true
	gameBoard.Cells[0][3].IsMine = true

	s, err := gameBoard.Reveal(0, 0)
	assert.NoError(t, err)
	assert.Equal(t, "1", s)
	assert.True(t, gameBoard.Cells[0][0].IsRevealed)
	assert.Equal(t, 1, gameBoard.Revealed)
	sFlagged, err := gameBoard.Reveal(0, 1)
	assert.Equal(t, "", sFlagged)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrIsFlagged.Error())
	assert.Equal(t, 1, gameBoard.Revealed)
	sRevealed, err := gameBoard.Reveal(0, 2)
	assert.Equal(t, "", sRevealed)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrIsAlreadyRevealed.Error())
	assert.Equal(t, 1, gameBoard.Revealed)
	sMine, err := gameBoard.Reveal(0, 3)
	assert.Equal(t, "*", sMine)
	assert.NoError(t, err)
	assert.Equal(t, 2, gameBoard.Revealed)

}

func TestRevealAdjacent(t *testing.T) {
	gameBoard := NewBoard(8, 16, 0, mockClickHandler)
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
func TestTableCell(t *testing.T) {
	cell := Cell{IsMine: false, IsFlag: false, IsRevealed: true, Adj: 1}
	tableCell := cell.TableCell()

	assert.NotNil(t, tableCell)
	assert.Equal(t, "[1]", tableCell.Text)
	assert.Equal(t, tview.AlignCenter, tableCell.Align)
	assert.Equal(t, cellColors["[1]"]["text"], tableCell.Color)
	assert.Equal(t, cellColors["[1]"]["bg"], tableCell.BackgroundColor)

	cell = Cell{IsMine: true, IsFlag: false, IsRevealed: true}
	tableCell = cell.TableCell()

	assert.NotNil(t, tableCell)
	assert.Equal(t, "[*]", tableCell.Text)
	assert.Equal(t, tview.AlignCenter, tableCell.Align)
	assert.Equal(t, cellColors["[*]"]["text"], tableCell.Color)
	assert.Equal(t, cellColors["[*]"]["bg"], tableCell.BackgroundColor)

	cell = Cell{IsMine: false, IsFlag: true, IsRevealed: false}
	tableCell = cell.TableCell()

	assert.NotNil(t, tableCell)
	assert.Equal(t, "[⚑]", tableCell.Text)
	assert.Equal(t, tview.AlignCenter, tableCell.Align)
	assert.Equal(t, cellColors["[⚑]"]["text"], tableCell.Color)
	assert.Equal(t, cellColors["[⚑]"]["bg"], tableCell.BackgroundColor)

	cell = Cell{IsMine: false, IsFlag: false, IsRevealed: false}
	tableCell = cell.TableCell()

	assert.NotNil(t, tableCell)
	assert.Equal(t, "[ ]", tableCell.Text)
	assert.Equal(t, tview.AlignCenter, tableCell.Align)
	assert.Equal(t, cellColors["[ ]"]["text"], tableCell.Color)
	assert.Equal(t, cellColors["[ ]"]["bg"], tableCell.BackgroundColor)
}
