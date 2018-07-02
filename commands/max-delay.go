package commands

import (
	"gopkg.in/urfave/cli.v1"
)

var MaxDelay = cli.Command{
	Name:   "max-delay",
	Usage:  "Cuts all delays between commands down to a given value",
	Action: maxDelayAction,
}

func maxDelayAction(c *cli.Context) (err error) {
	return
}
