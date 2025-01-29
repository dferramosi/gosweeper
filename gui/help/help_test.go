package help

import (
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestHelpItems(t *testing.T) {
	expectedHelpItems := [][]string{
		{"u/d/l/r arrows", "Select cells"},
		{"h,j,k,l", "Select cells"},
		{"enter/return, space", "Reveal cell"},
		{"f,z", "Flag cell"},
		{"?", "Show this help"},
		{"s", "Start a new game (not implemented)"},
		{"q/Ctrl+c", "Quit gosweeper"},
		{"ESC/?", "Close Help - Return to game"},
	}

	assert.Equal(t, expectedHelpItems, helpItems)
}

func TestUI(t *testing.T) {
	uiGrid := UI()
	assert.NotNil(t, uiGrid)
	assert.IsType(t, &tview.Grid{}, uiGrid)
}

func TestPopulateTable(t *testing.T) {
	tableWidget = tview.NewTable().SetBorders(false)
	populateTable()

	for i, row := range helpItems {
		for j, col := range row {
			cell := tableWidget.GetCell(i, j)
			assert.NotNil(t, cell)
			assert.Equal(t, col, cell.Text)
		}
	}
}
