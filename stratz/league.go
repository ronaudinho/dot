package stratz

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ronaudinho/dot/api"
)

func (c *Client) GetLeagueByID(id string) (api.League, error) {
	var l api.League
	return l, nil
}

func (c *Client) GetLeagues() ([]api.League, error) {
	url := fmt.Sprintf("%s/api/%s/league", c.baseURL, c.version)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var ls []api.League
	if err := json.Unmarshal(b, &ls); err != nil {
		return nil, err
	}
	return ls, nil
}

func (c *Client) GetLeagueMatches(id string) ([]api.Match, error) {
	url := fmt.Sprintf("%s/api/%s/league/%s/matches", c.baseURL, c.version, id)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var ms []api.Match
	if err := json.Unmarshal(b, &ms); err != nil {
		return nil, err
	}
	return ms, nil
}

func (c *Client) GetLeagueSeries(id string) ([]api.Match, error) {
	return nil, nil
}
