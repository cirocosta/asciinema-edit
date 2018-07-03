package commands

import (
	"gopkg.in/urfave/cli.v1"
)

var Speed = cli.Command{
	Name: "speed",
	Usage: `Updates the cast speed by a certain factor.

   If no file name is specified as a positional argument, a cast is
   expected to be serverd via stdin.

   Once the transformation has been performed, the resulting cast is
   either written to a file specified in the '--out' flag or to stdout
   (default).

EXAMPLES
   Make the whole cast ("123.cast") twice as fast:

     asciinema-edit speed --factor 2 ./123.cast

   Cut the speed in half:

     asciinema-edit speed --factor 0.5 ./123.cast

   Make only a certain part of the video twice as fast:

     asciinema-edit speed \
	--factor 2 \
	--start 12.231 \
	--factor 45.333 \
	./123.cast`,
	ArgsUsage: "[filename]",
	Action:    speedAction,
	Flags: []cli.Flag{
		cli.Float64Flag{
			Name:  "factor",
			Usage: "multiplying factor that will change the speed",
		},
		cli.Float64Flag{
			Name:  "start",
			Usage: "initial frame timestamp",
		},
		cli.Float64Flag{
			Name:  "end",
			Usage: "final frame timestamp",
		},
		cli.StringFlag{
			Name:  "out",
			Usage: "file to write the modified contents to",
		},
	},
}

func speedAction(c *cli.Context) (err error) {
	return
}
