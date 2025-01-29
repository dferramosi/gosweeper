package gui

import (
	"gosweeper/gui/game"
	"gosweeper/gui/help"
	"gosweeper/gui/settingsform"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var pages *tview.Pages
var activePage string

// Init starts the gui
func Init() {
	app = tview.NewApplication()
	pages = tview.NewPages()
	redraw := func() {
		app.Draw()
	}
	pages.AddPage("game", game.UI(redraw, switchToPage, gameEnd), true, false)
	pages.AddPage("help", help.UI(), true, false)
	pages.AddPage("settingsform", settingsform.UI(), true, true)
	activePage = "settingsform"

	app.SetInputCapture(eventHandler)
	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}
}

func gameEnd(win bool) {
	var msg string
	if win {
		msg = "You Win!"
	} else {
		msg = "You Lose!"
	}
	newGame := tview.NewModal()
	newGame.SetTitle(msg)
	newGame.SetText("New Game?").
		AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Yes" {
				switchToPage("settingsform")
			} else {
				switchToPage("quit")
			}
			pages.RemovePage("NewGameModal")
		})

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(newGame, 0, 1, true), 0, 1, true)
	pages.AddPage("NewGameModal", flex, true, true)
	newGame.SetFocus(0)
}

func switchToPage(page string) {
	activePage = page
	if page == "quit" {
		app.Stop()
	}
	pages.SwitchToPage(page)
}

func eventHandler(eventKey *tcell.EventKey) *tcell.EventKey {

	if eventKey.Rune() == 'q' {
		app.Stop()
		return nil
	}

	if activePage == "help" {
		return help.HandleEvents(eventKey, switchToPage)
	}

	if activePage == "game" {
		return game.HandleEvents(eventKey, switchToPage)
	}

	if activePage == "settingsform" {
		return settingsform.HandleEvents(eventKey, switchToPage)
	}

	return eventKey
}
