package game

import (
	"testing"

	"gosweeper/board"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

// func TestGameTimer(t *testing.T) {
// 	done = make(chan bool)
// 	ticker = time.NewTicker(time.Millisecond * 10)
// 	startTime = time.Now()

// 	redrawCalled := false
// 	redrawParent = func() {
// 		redrawCalled = true
// 	}

// 	go GameTimer()
// 	time.Sleep(time.Millisecond * 50)
// 	done <- true

// 	assert.True(t, redrawCalled)
// }

func TestUI(t *testing.T) {
	redraw := func() {}
	switchToPage := func(page string) {}
	gameEnd := func(win bool) {}

	uiGrid := UI(redraw, switchToPage, gameEnd)
	assert.NotNil(t, uiGrid)
	assert.IsType(t, &tview.Grid{}, uiGrid)
}

// func TestGameEnd(t *testing.T) {
// 	ticker = time.NewTicker(time.Second)
// 	done = make(chan bool)

// 	gameEndCalled := false
// 	gameEndParent = func(win bool) {
// 		gameEndCalled = true
// 	}

// 	go GameEnd(true)
// 	time.Sleep(time.Millisecond * 10)

// 	assert.True(t, gameEndCalled)
// }

// func TestReset(t *testing.T) {
// 	board.BoardRows = 10
// 	board.BoardCols = 10
// 	board.BoardMines = 10

// 	Reset()

// 	assert.NotNil(t, gameBoard)
// 	assert.Equal(t, board.BoardRows, gameBoard.Rows)
// 	assert.Equal(t, board.BoardCols, gameBoard.Columns)
// 	assert.Equal(t, board.BoardMines, gameBoard.Mines)
// 	assert.NotNil(t, ticker)
// 	assert.NotNil(t, done)
// }

func TestPopulateTable(t *testing.T) {
	board.BoardRows = 5
	board.BoardCols = 5
	board.BoardMines = 5
	gameBoard = board.NewBoard(board.BoardRows, board.BoardCols, board.BoardMines, clickHandler)

	tableWidget = tview.NewTable().SetBorders(true)
	populateTable(gameBoard)

	for row := 0; row < board.BoardRows; row++ {
		for col := 0; col < board.BoardCols; col++ {
			cell := tableWidget.GetCell(row, col)
			assert.NotNil(t, cell)
		}
	}
}

func TestFlagHandler(t *testing.T) {
	board.BoardRows = 5
	board.BoardCols = 5
	board.BoardMines = 5
	gameBoard = board.NewBoard(board.BoardRows, board.BoardCols, board.BoardMines, clickHandler)

	tableWidget = tview.NewTable().SetBorders(true)
	sideBar = tview.NewTable().SetBorders(false)
	tableWidget.Select(0, 0)

	flagHandler()

	cell := tableWidget.GetCell(0, 0)
	assert.NotNil(t, cell)
	assert.Equal(t, "[⚑]", cell.Text)
}

func TestSelectHandler(t *testing.T) {
	board.BoardRows = 5
	board.BoardCols = 5
	board.BoardMines = 0
	gameBoard = board.NewBoard(board.BoardRows, board.BoardCols, board.BoardMines, clickHandler)
	gameBoard.AddMine(5, 5)
	board.BoardMines = 1
	gameBoard.PopulateAdjacents()

	tableWidget = tview.NewTable().SetBorders(true)
	sideBar = tview.NewTable().SetBorders(false)

	selectHandler(0, 0)

	cell := tableWidget.GetCell(0, 0)
	assert.NotNil(t, cell)
	assert.Contains(t, cell.Text, "0")
}
