package stratz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ronaudinho/dot/api"
)

func (c *Client) GetLiveMatches() ([]api.Match, error) {
	query := `
{
  leagues(request: {betweenStartDateTime: 1633564800, betweenEndDateTime: 1633651200, skip: 0, take: 50, tiers: [INTERNATIONAL]}) {
    id
    name
    liveMatches {
      matchId
      radiantTeam {
        name
      }
      radiantScore
      direTeam {
        name
      }
      direScore
      radiantLead
      gameTime
      gameState
      players {
        hero {
          displayName
          shortName
        }
        name
        playerSlot
		isRadiant
      }
    }
  }
}
	`
	url := fmt.Sprintf("%s/graphql", c.baseURL)
	bod, _ := json.Marshal(struct {
		Query string `json:"query"`
	}{
		Query: query,
	})
	res, err := http.Post(url, "application/json", bytes.NewBuffer(bod))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r ResponseLive
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	// fuck type safety no?
	return toLive(r.Data.Leagues), nil
}

func toLive(leagues []League) []api.Match {
	var matches []api.Match
	for _, league := range leagues {
		for _, lm := range league.LiveMatches {
			var players []api.Player
			for _, p := range lm.Players {
				players = append(players, api.Player{
					IsRadiant: p.IsRadiant,
					Hero:      p.Hero.DisplayName,
				})
			}
			matches = append(matches, api.Match{
				ID:              lm.ID,
				RadiantName:     lm.RadiantTeam.Name,
				RadiantScore:    lm.RadiantScore,
				DireName:        lm.DireTeam.Name,
				DireScore:       lm.DireScore,
				RadiantLead:     lm.RadiantLead,
				DurationSeconds: lm.GameTime,
				Players:         players,
			})
		}
	}
	return matches
}
