package main

import (
	"github.com/ronaudinho/dot/stratz"
	"github.com/ronaudinho/dot/ui"
)

func main() {
	cli := stratz.NewDefaultClient()
	app := ui.NewApp(cli)
	if err := app.Main(); err != nil {
		panic(err)
	}
	if err := app.FE.Run(); err != nil {
		panic(err)
	}
}
