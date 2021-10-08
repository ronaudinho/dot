package api

type Service interface {
	LiveService
	LeagueService
	MatchService
}

type LiveService interface {
	GetLiveMatches() ([]Match, error)
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
	GetMatchDetail(string) (MatchDetail, error)
}

type Match struct {
	ID              int64 `json:"id"`
	RadiantID       int64 `json:"radiantTeamId"`
	DireID          int64 `json:"direTeamId"`
	RadiantName     string
	RadiantScore    int
	DireName        string
	DireScore       int
	RadiantLead     int
	StartDateTime   int64    `json:"startDateTime"`
	EndDateTime     int64    `json:"endDateTime"`
	DurationSeconds int      `json:"durationSeconds"`
	DidRadiantWin   bool     `json:"didRadiantWin"`
	Players         []Player `json:"players"`
}

type Player struct {
	IsRadiant           bool `json:"isRadiant"`
	Hero                string
	ExperiencePerMinute int `json:"experiencePerMinute"`
	NetWorth            int `json:"networth"`
	Assists             int `json:"numAssists"`
	Deaths              int `json:"numDeaths"`
	Kills               int `json:"numKills"`
}

type MatchDetail struct {
	// placeholder
}

type League struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	DisplayName    string `json:"displayName"`
	HasLiveMatches bool   `json:"hasLiveMatches"`
}
