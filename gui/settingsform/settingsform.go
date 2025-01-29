package settingsform

import (
	"fmt"
	"gosweeper/board"
	"gosweeper/gui/game"
	"gosweeper/logger"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var command string
var uiScreen *tview.Grid
var uiForm *tview.Form

// UI creates the help window
func UI() *tview.Form {
	uiForm = tview.NewForm().
		AddInputField("Rows", fmt.Sprint(board.BoardRows), 3, isInt, handleRowChange).
		AddInputField("Columns", fmt.Sprint(board.BoardCols), 3, isInt, handleColChange).
		AddInputField("Mines", fmt.Sprint(board.BoardMines), 3, isInt, handleMineChange).
		AddButton("Start", handlePressStart).
		AddButton("Quit", handlePressQuit)

	uiForm.SetFieldBackgroundColor(tcell.ColorGold).SetFieldTextColor(tcell.ColorBlack)
	uiForm.SetBorder(true).SetTitle("Settings")

	return uiForm
}

func isInt(s string, lc rune) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func handleRowChange(rows string) {
	logger.DebugLogf("Rows changed to %s", rows)
	i, _ := strconv.Atoi(rows)
	board.BoardRows = i
}

func handleColChange(cols string) {
	logger.DebugLogf("Cols changed to %s", cols)
	i, _ := strconv.Atoi(cols)
	board.BoardCols = i
}

func handleMineChange(mines string) {
	logger.DebugLogf("Mines changed to %s", mines)
	i, _ := strconv.Atoi(mines)
	board.BoardMines = i
}

func handlePressStart() {
	game.Reset()
	parentScreenSwitchPage("game")
}

func handlePressQuit() {
	parentScreenSwitchPage("quit")
}
