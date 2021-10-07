package ui

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

func (a *App) Main() error {
	ls, err := a.be.GetLeagues()
	if err != nil {
		return err
	}

	menu := tview.NewList()
	for i, l := range ls {
		var name string
		if l.DisplayName != "" {
			name = l.DisplayName
		} else {
			name = l.Name
		}
		page := fmt.Sprintf("league/%d", l.ID)
		menu.AddItem(name, "", rune(i), func() {
			a.pages.SwitchToPage(page)
		})
		a.League(strconv.FormatInt(l.ID, 10))
	}
	menu.AddItem("quit", "", 'q', func() {
		a.FE.Stop()
	})

	a.pages.AddPage("menu", menu, true, true)
	a.FE.SetRoot(a.pages, true)
	return nil
}
