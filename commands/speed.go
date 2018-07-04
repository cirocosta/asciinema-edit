package commands

import (
	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/commands/transformer"
	"github.com/cirocosta/asciinema-edit/editor"
	"gopkg.in/urfave/cli.v1"
)

var Speed = cli.Command{
	Name: "speed",
	Usage: `Updates the cast speed by a certain factor.

   If no file name is specified as a positional argument, a cast is
   expected to be serverd via stdin.

   If no range is specified (start=0, end=0), the whole event stream
   is processed.

   Once the transformation has been performed, the resulting cast is
   either written to a file specified in the '--out' flag or to stdout
   (default).

EXAMPLES:
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
			Usage: "number by which delays are multiplied by",
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

type speedTransformation struct {
	from   float64
	to     float64
	factor float64
}

func (t *speedTransformation) Transform(c *cast.Cast) (err error) {
	if t.from == 0 && t.to == 0 {
		t.from = c.EventStream[0].Time
		t.to = c.EventStream[len(c.EventStream)-1].Time
	}

	err = editor.Speed(c, t.factor, t.from, t.to)
	return
}

func speedAction(c *cli.Context) (err error) {
	var (
		input          = c.Args().First()
		output         = c.String("out")
		transformation = &speedTransformation{
			factor: c.Float64("factor"),
			from:   c.Float64("start"),
			to:     c.Float64("end"),
		}
	)

	t, err := transformer.New(transformation, input, output)
	if err != nil {
		err = cli.NewExitError(err, 1)
		return
	}

	err = t.Transform()
	if err != nil {
		err = cli.NewExitError(err, 1)
		return
	}

	t.Close()
	return
}
