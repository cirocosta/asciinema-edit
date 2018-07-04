package commands

import (
	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/commands/transformer"
	"github.com/cirocosta/asciinema-edit/editor"
	"gopkg.in/urfave/cli.v1"
)

var Cut = cli.Command{
	Name: "cut",
	Usage: `Removes a certain range of time frames.

   If no file name is specified as a positional argument, a cast is
   expected to be served via stdin.

   Once the transformation has been performed, the resulting cast is
   either written to a file specified in the '--out' flag or to stdout
   (default).

EXAMPLES:
   Remove frames from 12.2s to 16.3s from the cast passed in the commands
   stdin.

     cat 1234.cast | \
       asciinema-edit cut \
         --from=12.2 --to=15.3

   Remove the exact frame at timestamp 12.2 from the cast file named
   1234.cast.

     asciinema-edit cut \
       --from=12.2 --to=12.2 \
       1234.cast`,
	ArgsUsage: "[filename]",
	Action:    cutAction,
	Flags: []cli.Flag{
		cli.Float64Flag{
			Name:  "start",
			Usage: "initial frame timestamp (required)",
		},
		cli.Float64Flag{
			Name:  "end",
			Usage: "final frame timestamp (required)",
		},
		cli.StringFlag{
			Name:  "out",
			Usage: "file to write the modified contents to",
		},
	},
}

type cutTransformation struct {
	from float64
	to   float64
}

func (t *cutTransformation) Transform(c *cast.Cast) (err error) {
	err = editor.Cut(c, t.from, t.to)
	return
}

func cutAction(c *cli.Context) (err error) {
	var (
		input          = c.Args().First()
		output         = c.String("out")
		transformation = &cutTransformation{
			from: c.Float64("start"),
			to:   c.Float64("end"),
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
