package ui

import (
	"github.com/ronaudinho/dot/api"

	"github.com/rivo/tview"
)

type App struct {
	FE    *tview.Application
	pages *tview.Pages
	be    api.Service
}

func NewApp(svc api.Service) *App {
	return &App{
		FE:    tview.NewApplication(),
		pages: tview.NewPages(),
		be:    svc,
	}
}
