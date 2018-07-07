package commands

import (
	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/commands/transformer"
	"gopkg.in/urfave/cli.v1"

	_ "github.com/cirocosta/asciinema-edit/editor"
)

var Convert = cli.Command{
	Name: "convert",
	Usage: `Converts asciinema casts from one format to another.

   If no file name is specified as a positional argument, a cast is
   expected to be served via stdin.

EXAMPLES:
   Convert a cast (123.cast) from V1 format to V2 format:

       asciinema-edit convert --from=v1 --to=v2 123.cast

   Convert a cast (321.cast) from V2 format to V1 format:

       asciinema-edit convert --from=v2 --to=v1 321.cast`,
	ArgsUsage: "[filename]",
	Action:    cutAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "from",
			Usage: "format of the source cast(v1 or v2)",
		},
		cli.StringFlag{
			Name:  "to",
			Usage: "format of the destination cast (v1 or v2)",
		},
		cli.StringFlag{
			Name:  "out",
			Usage: "file to write the modified contents to",
		},
	},
}

type convertTransformation struct {
	from string
	to   string
}

// TODO probably `cast` should now support type casting somehow?
func (t *convertTransformation) Transform(c *cast.Cast) (err error) {
	// err = editor.Convert(c, t.from, t.to)
	return
}

func convertAction(c *cli.Context) (err error) {
	var (
		input          = c.Args().First()
		output         = c.String("out")
		transformation = &convertTransformation{
			from: c.String("from"),
			to:   c.String("to"),
		}
	)

	t, err := transformer.New(transformation, input, output)
	if err != nil {
		err = cli.NewExitError(err, 1)
		return
	}
	defer t.Close()

	err = t.Transform()
	if err != nil {
		err = cli.NewExitError(err, 1)
		return
	}

	return
}
