package commands

import (
	"gopkg.in/urfave/cli.v1"
)

var Cut = cli.Command{
	Name:   "cut",
	Usage:  "Removes a certain range of time frames",
	Action: cutAction,
}

func cutAction(c *cli.Context) (err error) {
	return
}
