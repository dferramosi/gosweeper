package game

import (
	"fmt"
	"gosweeper/board"
	"gosweeper/logger"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var redrawParent func()
var switchToPageParent func(string)
var gameEndParent func(bool)
var uiScreen *tview.Grid
var tableWidget *tview.Table
var sideBar *tview.Table
var gameBoard *board.GameBoard
var ticker *time.Ticker
var startTime time.Time
var done chan bool
var clickFlag bool

func GameTimer() {
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			logger.DebugLogf("Tick at %v", t)
			diff := t.Sub(startTime).Seconds()
			sideBar.SetCellSimple(3, 2, fmt.Sprintf("%03d", int(diff)))
			redrawParent()
		}
	}
}

// UI creates the game window
func UI(redraw func(), switchToPage func(string), gameEnd func(bool)) *tview.Grid {
	redrawParent = redraw
	switchToPageParent = switchToPage
	gameEndParent = gameEnd

	uiScreen = tview.NewGrid().SetRows(0).SetColumns(-4, -1).SetBorders(true)

	tableWidget = tview.NewTable().SetBorders(true)
	tableWidget.SetSelectedFunc(selectHandler)

	tableWidget.SetSelectable(true, true)

	uiFrame := tview.NewFrame(tableWidget)

	uiScreen.AddItem(uiFrame, 0, 0, 1, 1, 0, 0, true)

	sideBar = tview.NewTable().SetBorders(false)
	sideBar.SetSelectable(false, false)
	sideBar.SetCellSimple(2, 0, "Mines Left")
	sideBar.SetCellSimple(2, 2, fmt.Sprint(0))
	sideBar.SetCellSimple(3, 0, "Time")
	sideBar.SetCellSimple(4, 0, "Click Mode")
	sideBar.SetCellSimple(4, 2, "Flag")
	ticker = time.NewTicker(time.Second)
	done = make(chan bool)
	sideBarFrame := tview.NewFrame(sideBar).
		AddText("GoSweeper", true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText("Enter/Space to reveal,", false, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText("z/f to flag,", false, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText("? for Help", false, tview.AlignCenter, tcell.ColorDarkMagenta)
	uiScreen.AddItem(sideBarFrame, 0, 1, 1, 1, 0, 0, false)

	return uiScreen
}

func flipClickMode() {
	if clickFlag {
		clickFlag = false
		sideBar.SetCellSimple(4, 2, "Select")
	} else {
		clickFlag = true
		sideBar.SetCellSimple(4, 2, "Flag")
	}
	redrawParent()
}

func GameEnd(win bool) {
	ticker.Stop()
	done <- true
	gameEndParent(win)
}

func Reset() {
	logger.DebugLogf("Creating new game board with %d rows, %d cols, %d mines", board.BoardRows, board.BoardCols, board.BoardMines)
	clickFlag = false

	gameBoard = board.NewBoard(board.BoardRows, board.BoardCols, board.BoardMines, clickHandler)
	populateTable(gameBoard)
	sideBar.SetCellSimple(2, 2, fmt.Sprint(gameBoard.Mines))
	startTime = time.Now()
	ticker.Reset(time.Second)
	done = make(chan bool)
	go GameTimer()
	logger.DebugLogf("%v", done)
}

func populateTable(gameBoard *board.GameBoard) {
	tableWidget.Clear()
	for row := 0; row < board.BoardRows; row++ {
		for col := 0; col < board.BoardCols; col++ {
			c, _ := gameBoard.GetCell(row, col)
			tableWidget.SetCell(row, col, c.TableCell().SetExpansion(1))
		}
	}
}

func clickHandler() bool {
	row, col := tableWidget.GetSelection()
	logger.DebugLogf("Clicking Cell(%d,%d)", row, col)
	if clickFlag {
		flagHandler()
		return true
	}
	return false
}

func flagHandler() {
	row, col := tableWidget.GetSelection()
	logger.DebugLogf("Flagging Cell(%d,%d)", row, col)
	newVal := gameBoard.Flag(row, col)
	logger.DebugLogf("Flagged Cell(%d,%d): %v", row, col, newVal)
	c, _ := gameBoard.GetCell(row, col)
	tableWidget.SetCell(row, col, c.TableCell())
	sideBar.SetCellSimple(2, 2, fmt.Sprint(gameBoard.Mines-gameBoard.Flags))
}

func selectHandler(row, col int) {
	cell, _ := gameBoard.GetCell(row, col)
	logger.DebugLogf("Selecting Cell(%d,%d): %v", row, col, cell.String())
	newVal, err := gameBoard.Reveal(row, col)
	if err != nil {
		logger.DebugLogf("Error revealing cell: %v", err)
		return
	}
	logger.DebugLogf("Revealed Cell(%d,%d): %v", row, col, newVal)
	if newVal == "0" {
		logger.DebugLogf("Revealing adjacent cells")
		gameBoard.RevealAdjacent(row, col, []string{}, true)
		populateTable(gameBoard)
	} else {
		tableWidget.SetCell(row, col, cell.TableCell())
	}
	if cell.IsMine {
		logger.DebugLogf("Game Over")
		GameEnd(false)
	}
	if gameBoard.Revealed == (board.BoardRows*board.BoardCols - board.BoardMines) {
		logger.DebugLogf("Game Win")
		GameEnd(true)
	}
}
