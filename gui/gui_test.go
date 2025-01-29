package gui

// import (
// 	"testing"

// 	"github.com/gdamore/tcell/v2"
// 	"github.com/rivo/tview"
// 	"github.com/stretchr/testify/assert"
// )

// unsure how to test sending an an input event to a modal

// func TestGameEnd(t *testing.T) {
// 	app = tview.NewApplication()
// 	pages = tview.NewPages()

// 	// Test for win scenario
// 	gameEnd(true)
// 	pageName, _ := pages.GetFrontPage()
// 	assert.Equal(t, "NewGameModal", pageName)

// 	// Test for lose scenario
// 	gameEnd(false)
// 	pageName, modalP := pages.GetFrontPage()
// 	assert.Equal(t, "NewGameModal", pageName)

// 	// Simulate button press "Yes"
// 	modal := modalP.(*tview.Flex).GetItem(0).(*tview.Flex).GetItem(0).(*tview.Modal)
// 	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
// 		assert.Equal(t, "Yes", buttonLabel)
// 		switchToPage("settingsform")
// 	})
// 	modal.InputHandler()(tcell.NewEventKey(tcell.KeyEnter, '\n', tcell.ModNone), nil)

// 	assert.Equal(t, "settingsform", activePage)

// 	// Simulate button press "No"
// 	gameEnd(false)
// 	modal = modalP.(*tview.Flex).GetItem(0).(*tview.Flex).GetItem(0).(*tview.Modal)
// 	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
// 		assert.Equal(t, "No", buttonLabel)
// 		switchToPage("quit")
// 	})
// 	modal.InputHandler()(tcell.NewEventKey(tcell.KeyEnter, ' ', tcell.ModNone), nil)

// 	assert.Equal(t, "quit", activePage)
// }
