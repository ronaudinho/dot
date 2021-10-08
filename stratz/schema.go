package stratz

type ResponseLive struct {
	Data struct {
		Leagues []League `json:"leagues"`
	} `json:"data"`
}

type League struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	LiveMatches []LiveMatch `json:"liveMatches"`
}

type LiveMatch struct {
	ID           int64    `json:"matchId"`
	RadiantTeam  Team     `json:"radiantTeam"`
	DireTeam     Team     `json:"direTeam"`
	RadiantScore int      `json:"radiantScore"`
	DireScore    int      `json:"direScore"`
	RadiantLead  int      `json:"radiantLead"`
	GameTime     int      `json:"gameTime"`
	GameState    string   `json:"gameState"` // enum
	Players      []Player `json:"players"`
}

type Team struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// LivePlayer instead
type Player struct {
	Name       string `json:"name"`
	Hero       Hero   `json:"hero"`
	PlayerSlot int    `json:"playerSlot"`
	IsRadiant  bool   `json:"isRadiant"`
}

type Hero struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	ShortName   string `json:"shortName"`
}
