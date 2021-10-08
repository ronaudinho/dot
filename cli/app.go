package cli

import (
	"github.com/ronaudinho/dot/api"
)

type App struct {
	// printer
	be api.Service
}

func NewApp(svc api.Service) *App {
	return &App{
		be: svc,
	}
}
