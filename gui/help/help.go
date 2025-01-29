package help

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var helpItems = [][]string{
	{"u/d/l/r arrows", "Select cells"},
	{"h,j,k,l", "Select cells"},
	{"enter/return, space", "Reveal cell"},
	{"f,z", "Flag cell"},
	{"?", "Show this help"},
	{"s", "Start a new game (not implemented)"},
	{"q/Ctrl+c", "Quit gosweeper"},
	{"ESC/?", "Close Help - Return to game"},
}

var uiScreen *tview.Grid
var tableWidget *tview.Table

// UI creates the help window
func UI() *tview.Grid {
	uiScreen = tview.NewGrid().SetRows(0).SetColumns(0).SetBorders(true)
	tableWidget = tview.NewTable().SetBorders(false)

	populateTable()

	uiFrame := tview.NewFrame(tableWidget).
		AddText("Help", true, tview.AlignCenter, tcell.ColorDarkMagenta).
		AddText("ESC or F1 to leave Help", false, tview.AlignCenter, tcell.ColorDarkMagenta)

	uiScreen.AddItem(uiFrame, 0, 0, 1, 1, 0, 0, false)

	return uiScreen
}

func populateTable() {
	tableWidget.Clear()
	for i := range helpItems {
		for j, col := range helpItems[i] {
			tableWidget.SetCell(i, j, tview.NewTableCell(col).SetExpansion(1))
		}
	}
}
