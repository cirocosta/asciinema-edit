package commands

import (
	"gopkg.in/urfave/cli.v1"
)

var MaxDelay = cli.Command{
	Name:   "max-delay",
	Usage:  "cuts all delays between commands up to a maximum value",
	Action: maxDelayAction,
}

func maxDelayAction(c *cli.Context) (err error) {
	return
}
