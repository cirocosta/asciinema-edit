package commands

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/commands/transformer"
	"github.com/cirocosta/asciinema-edit/editor"
	"gopkg.in/urfave/cli.v1"
)

var Quantize = cli.Command{
	Name: "quantize",
	Usage: `Updates the cast delays following quantization ranges.

   The command acts on the delays between the frames, reducing such
   timings to the lowest value defined in a given range that they
   lie in.

   For instance, consider the following timestamps:

      1  2  5  9 10 11

   Assuming that we quantize over [2,6), we'd cut any delays between 2 and
   6 seconds to 2 second:

      1  2  4  6  7  8

   This can be more easily visualized by looking at the delay quantization:

      delta = 1.000000 | qdelta = 1.000000
      delta = 3.000000 | qdelta = 2.000000
      delta = 4.000000 | qdelta = 2.000000
      delta = 1.000000 | qdelta = 1.000000
      delta = 1.000000 | qdelta = 1.000000

   If no file name is specified as a positional argument, a cast is
   expected to be serverd via stdin.

   Once the transformation has been performed, the resulting cast is
   either written to a file specified in the '--out' flag or to stdout
   (default).

EXAMPLES:
   Make the whole cast have a maximum delay of 1s:

     asciinema-edit quantize --range 2 ./123.cast

   Make the whole cast have time delays between 300ms and 1s cut to
   300ms, delays between 1s and 2s cut to 1s and any delays bigger
   than 2s, cut down to 2s:

     asciinema-edit quantize \
       --range 0.3,1 \
       --range 1,2 \
       --range 2 \
       ./123.cast`,
	ArgsUsage: "[filename]",
	Action:    quantizeAction,
	Flags: []cli.Flag{
		cli.StringSliceFlag{
			Name:  "range",
			Usage: "quantization ranges (comma delimited)",
		},
		cli.StringFlag{
			Name:  "out",
			Usage: "file to write the modified contents to",
		},
	},
}

type quantizeTransformation struct {
	ranges []editor.QuantizeRange
}

func (t *quantizeTransformation) Transform(c *cast.Cast) (err error) {
	err = editor.Quantize(c, t.ranges)
	return
}

func parseQuantizeRanges(ranges []string) (res []editor.QuantizeRange, err error) {
	res = make([]editor.QuantizeRange, len(ranges))

	for idx, rangeStr := range ranges {
		cols := strings.Split(rangeStr, ",")

		if len(cols) > 2 {
			err = errors.Errorf(
				"invalid range format: must be `value[,value]`")
			return
		}

		if len(cols) == 2 {
			res[idx].To, err = strconv.ParseFloat(cols[1], 64)
			if err != nil {
				err = errors.Errorf(
					"malformed range: not a float %s", cols[1])
				return
			}
		}

		res[idx].From, err = strconv.ParseFloat(cols[0], 64)
		if err != nil {
			err = errors.Errorf(
				"malformed range: not a float %s", cols[0])
			return
		}
	}

	return
}

func quantizeAction(c *cli.Context) (err error) {
	var (
		input          = c.Args().First()
		output         = c.String("out")
		ranges         = c.StringSlice("range")
		transformation = &quantizeTransformation{}
	)

	if len(ranges) == 0 {
		err = cli.NewExitError("a range must be specified.", 1)
		return
	}

	transformation.ranges, err = parseQuantizeRanges(ranges)
	if err != nil {
		err = cli.NewExitError(err, 1)
		return
	}

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
