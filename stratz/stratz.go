package stratz

const (
	DefaultBaseURL = "https://api.stratz.com"
	DefaultVersion = "v1"
	// api/v1/match/6210663857/breakdown"
)

type Client struct {
	baseURL string
	version string
}

func NewDefaultClient() *Client {
	return &Client{
		baseURL: DefaultBaseURL,
		version: DefaultVersion,
	}
}
