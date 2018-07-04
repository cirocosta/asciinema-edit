package transformer

import (
	"os"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/pkg/errors"
)

type Transformation interface {
	Transform(c *cast.Cast) (err error)
}

type Transformer struct {
	input          *os.File
	output         *os.File
	cast           *cast.Cast
	Transformation Transformation
}

func New(t Transformation, input, output string) (m *Transformer, err error) {
	if t == nil {
		err = errors.Errorf("a transformation must be specified")
		return
	}

	m = &Transformer{
		input:          os.Stdin,
		output:         os.Stdout,
		Transformation: t,
	}

	if input != "" {
		m.input, err = os.Open(input)
		if err != nil {
			err = errors.Wrapf(err,
				"failed to open input file %s", input)
			return
		}

		var stat os.FileInfo
		stat, err = m.input.Stat()
		if err != nil {
			m.input.Close()
			err = errors.Wrapf(err,
				"failed to retrieve info about input file %s", input)
			return
		}

		if stat.IsDir() {
			err = errors.Errorf("input file must not be a directory")
			return
		}
	}

	if output != "" {
		m.output, err = os.Create(output)
		if err != nil {
			m.input.Close()
			err = errors.Wrapf(err,
				"failed to open output file %s", output)
			return
		}
	}

	return
}

func (m *Transformer) Apply() (err error) {
	err = m.Transformation.Transform(m.cast)
	return
}

func (m *Transformer) Close() (err error) {
	if m.output != nil && m.output != os.Stdout {
		m.output.Close()
	}

	if m.input != nil && m.input != os.Stdin {
		m.input.Close()
	}

	return
}
