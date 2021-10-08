package stratz

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ronaudinho/dot/api"
)

func (c *Client) GetMatchByID(id string) (api.Match, error) {
	var m api.Match
	return m, nil
}

func (c *Client) GetMatches() ([]api.Match, error) {
	return nil, nil
}

func (c *Client) GetMatchDetail(id string) (api.MatchDetail, error) {
	var m api.MatchDetail
	url := fmt.Sprintf("%s/api/%s/match/%s/live", c.baseURL, c.version, id)

	res, err := http.Get(url)
	if err != nil {
		return m, err
	}
	defer res.Body.Close()

	// if live match has ended, stratz returns 204
	if res.StatusCode == http.StatusNoContent {
		return c.GetMatchBreakdown(id)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return m, err
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return m, err
	}
	return m, nil
}

func (c *Client) GetMatchBreakdown(id string) (api.MatchDetail, error) {
	var m api.MatchDetail
	url := fmt.Sprintf("%s/api/%s/match/%s/breakdown", c.baseURL, c.version, id)

	res, err := http.Get(url)
	if err != nil {
		return m, err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return m, err
	}
	if err := json.Unmarshal(b, &m); err != nil {
		return m, err
	}
	return m, nil
}
