package settingsform

import (
	"testing"

	"gosweeper/board"

	"github.com/stretchr/testify/assert"
)

func TestIsInt(t *testing.T) {
	assert.True(t, isInt("123", '1'))
	assert.False(t, isInt("abc", 'a'))
}

func TestHandleRowChange(t *testing.T) {
	initialRows := board.BoardRows
	handleRowChange("10")
	assert.Equal(t, 10, board.BoardRows)
	board.BoardRows = initialRows // Reset to initial value
}

func TestHandleColChange(t *testing.T) {
	initialCols := board.BoardCols
	handleColChange("15")
	assert.Equal(t, 15, board.BoardCols)
	board.BoardCols = initialCols // Reset to initial value
}

func TestHandleMineChange(t *testing.T) {
	initialMines := board.BoardMines
	handleMineChange("20")
	assert.Equal(t, 20, board.BoardMines)
	board.BoardMines = initialMines // Reset to initial value
}

func TestHandlePressQuit(t *testing.T) {
	// Mock the parentScreenSwitchPage function
	originalSwitchPage := parentScreenSwitchPage
	defer func() { parentScreenSwitchPage = originalSwitchPage }()
	parentScreenSwitchPage = func(page string) {
		assert.Equal(t, "quit", page)
	}

	handlePressQuit()
}
