package api

type Service interface {
	LeagueService
	MatchService
}

type LeagueService interface {
	GetLeagueByID(string) (League, error)
	GetLeagues() ([]League, error)
	GetLeagueMatches(string) ([]Match, error)
	GetLeagueSeries(string) ([]Match, error)
}

type MatchService interface {
	GetMatchByID(string) (Match, error)
	GetMatches() ([]Match, error)
	GetMatchLive(string) (Match, error)
}

type Match struct {
	ID              int64 `json:"id"`
	RadiantID       int64 `json:"radiantTeamId"`
	DireID          int64 `json:"direTeamId"`
	StartDateTime   int64 `json:"startDateTime"`
	EndDateTime     int64 `json:"endDateTime"`
	DurationSeconds int   `json:"durationSeconds"`
	DidRadiantWin   bool  `json:"didRadiantWin"`
	Players         []struct {
		IsRadiant           bool `json:"isRadiant"`
		HeroID              int  `json:"heroId"`
		ExperiencePerMinute int  `json:"experiencePerMinute"`
		NetWorth            int  `json:"networth"`
		Assists             int  `json:"numAssists"`
		Deaths              int  `json:"numDeaths"`
		Kills               int  `json:"numKills"`
	} `json:"players"`
}

type League struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	DisplayName    string `json:"displayName"`
	HasLiveMatches bool   `json:"hasLiveMatches"`
}
