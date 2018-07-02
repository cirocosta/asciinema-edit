package main

import (
	"gopkg.in/urfave/cli.v1"
)

var (
	version = "dev"
)

func main () {
	app := cli.NewApp()

	app.Commands = []cli.Command{}
	app.Version = version
}
