package ui

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

func (a *App) League(id string) error {
	ms, err := a.be.GetLeagueMatches(id)
	if err != nil {
		return err
	}

	menu := tview.NewList()
	for i, m := range ms {
		var rk, dk, nw, xp int
		var rh, dh []string
		for _, p := range m.Players {
			if p.IsRadiant {
				rk += p.Kills
				nw += p.NetWorth
				xp += (p.ExperiencePerMinute * m.DurationSeconds / 60)
			} else {
				dk += p.Kills
				nw -= p.NetWorth
				xp -= (p.ExperiencePerMinute * m.DurationSeconds / 60)
			}
		}
		name := fmt.Sprintf("%d %d - %d %d", m.RadiantID, rk, dk, m.DireID) // convert to name
		desc := fmt.Sprintf("%s net: %d exp: %d %s", strings.Join(rh, ","), nw, xp, strings.Join(dh, ","))
		page := fmt.Sprintf("match/%d", m.ID)
		menu.AddItem(name, desc, rune(i), func() {
			a.pages.SwitchToPage(page)
		})
	}
	menu.AddItem("back", "", 'b', func() {
		a.pages.SwitchToPage("menu")
	})

	a.pages.AddPage(fmt.Sprintf("league/%s", id), menu, true, true)
	return nil
}
