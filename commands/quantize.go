package commands

import (
	"gopkg.in/urfave/cli.v1"
)

var Quantize = cli.Command{
	Name:   "quantize",
	Action: quantizeAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "range",
		},
	},
}

func quantizeAction(c *cli.Context) (err error) {
	return
}
