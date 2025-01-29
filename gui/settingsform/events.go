package settingsform

import "github.com/gdamore/tcell/v2"

var parentScreenSwitchPage func(string)

// HandleEvents dispatches key events for the help package
func HandleEvents(eventKey *tcell.EventKey, switchToPage func(string)) *tcell.EventKey {
	parentScreenSwitchPage = switchToPage

	return eventKey
}
