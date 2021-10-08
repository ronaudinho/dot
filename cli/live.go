package cli

import (
	"fmt"
	"strings"
)

func (a *App) Live() {
	matches, err := a.be.GetLiveMatches()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, m := range matches {
		var rh, dh []string
		for _, p := range m.Players {
			if p.IsRadiant {
				rh = append(rh, p.Hero)
			} else {
				dh = append(dh, p.Hero)
			}
		}
		fmt.Printf("* %s %d - %d %s\n", m.RadiantName, m.RadiantScore, m.DireScore, m.DireName)
		fmt.Printf(" %s net: %d %s\n", strings.Join(rh, ","), m.RadiantLead, strings.Join(dh, ","))
	}
}
