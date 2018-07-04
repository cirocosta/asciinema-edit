package main

import (
	"os"

	"github.com/cirocosta/asciinema-edit/commands"
	"gopkg.in/urfave/cli.v1"
)

var (
	version = "dev"
	commit  = "HEAD"
)

func main() {
	app := cli.NewApp()

	app.Version = version + " - " + commit
	app.Usage = "edit recorded asciinema casts"
	app.Description = `asciinema-edit provides missing features from the "asciinema" tool
   when it comes to editing a cast that has already been recorded.`
	app.Commands = []cli.Command{
		commands.Cut,
		commands.Speed,
	}

	app.Run(os.Args)
}
