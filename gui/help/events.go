package help

import "github.com/gdamore/tcell/v2"

// HandleEvents dispatches key events for the help package
func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	if eventKey.Key() == tcell.KeyEsc || eventKey.Rune() == '?' {
		switchToPage("game")
		return eventKey
	}

	return eventKey
}
