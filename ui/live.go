package ui

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

func (a *App) Live() (*tview.List, error) {
	matches, err := a.be.GetLiveMatches()
	if err != nil {
		return nil, err
	}

	list := tview.NewList()
	for i, m := range matches {
		var rh, dh []string
		for _, p := range m.Players {
			if p.IsRadiant {
				rh = append(rh, p.Hero)
			} else {
				dh = append(dh, p.Hero)
			}
		}
		name := fmt.Sprintf("%s %d - %d %s", m.RadiantName, m.RadiantScore, m.DireScore, m.DireName)
		desc := fmt.Sprintf("%s net: %d %s", strings.Join(rh, ","), m.RadiantLead, strings.Join(dh, ","))
		page := fmt.Sprintf("match/%d", m.ID)
		list.AddItem(name, desc, rune(i), func() {
			a.pages.SwitchToPage(page)
		})
	}

	return list, nil
}
