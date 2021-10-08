package ui

import (
	"github.com/rivo/tview"
)

// Main draws the main menu of the UI
func (a *App) Main() error {
	menu := tview.NewList()

	lives, err := a.Live()
	if err != nil {
		//
	}
	if lives != nil {
		menu = lives
	}

	menu.AddItem("quit", "", 'q', func() {
		a.FE.Stop()
	})

	a.pages.AddPage("menu", menu, true, true)
	a.FE.SetRoot(a.pages, true)
	return nil
}
