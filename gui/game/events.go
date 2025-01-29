package game

import "github.com/gdamore/tcell/v2"

// HandleEvents dispatches key events for the game package
func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	switch eventKey.Rune() {
	case '?':
		switchToPage("help")
	case 'f', 'z':
		flagHandler()
	case ' ':
		selectHandler(tableWidget.GetSelection())
	}

	switch eventKey.Key() {
	case tcell.KeyTab:
		// flag
	}

	return eventKey
}
