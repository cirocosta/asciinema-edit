package commands

import (
	"os"

	"github.com/cirocosta/asciinema-edit/cast"
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

EXAMPLES
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

func cutAction(c *cli.Context) (err error) {
	var (
		from     = c.Float64("start")
		to       = c.Float64("end")
		filename = c.Args().First()
		out      = c.String("out")

		inputFile  *os.File = os.Stdin
		outputFile *os.File = os.Stdout
	)

	if filename != "" {
		inputFile, err = os.Open(filename)
		if err != nil {
			err = cli.NewExitError(
				"failed to open file "+filename+
					" - "+err.Error(), 1)
			return
		}
		defer inputFile.Close()
	}

	if out != "" {
		outputFile, err = os.Create(out)
		if err != nil {
			err = cli.NewExitError(
				"failed to create and open output file "+out+
					" - "+err.Error(), 1)
			return
		}
		defer outputFile.Close()

	}

	decodedCast, err := cast.Decode(inputFile)
	if err != nil {
		err = cli.NewExitError(
			"failed to decode contents as asciinema cast (v2) - "+
				err.Error(), 1)
		return
	}

	err = editor.Cut(decodedCast, from, to)
	if err != nil {
		err = cli.NewExitError(
			"failed to cut cast - "+err.Error(), 1)
		return
	}

	err = cast.Encode(outputFile, decodedCast)
	if err != nil {
		err = cli.NewExitError(
			"failed to save modified cast to file "+out+
				" - "+err.Error(), 1)
		return
	}

	return
}
