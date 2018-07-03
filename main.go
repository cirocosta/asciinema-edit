package main

import (
	"os"

	"github.com/cirocosta/asciinema-edit/commands"
	"gopkg.in/urfave/cli.v1"
)

var (
	version = "dev"
)

func main() {
	app := cli.NewApp()

	app.Version = version
	app.Usage = "edit recorded asciinema casts"
	app.Description = `asciinema-edit provides missing features from the "asciinema" tool
   when it comes to editing a cast that has already been recorded.`
	app.Commands = []cli.Command{
		commands.AddDelay,
		commands.Cut,
		commands.MaxDelay,
	}

	app.Run(os.Args)
}
