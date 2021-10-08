package main

import (
	"github.com/ronaudinho/dot/cli"
	"github.com/ronaudinho/dot/stratz"
)

func main() {
	api := stratz.NewDefaultClient()
	app := cli.NewApp(api)
	app.Live()
}
