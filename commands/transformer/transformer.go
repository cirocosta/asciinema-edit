// Package transformer is an utility package designed to aid
// the process of taking a cast from a given input, applying
// a transformation to it and then outputting the transformed
// cast to somewhere.
package transformer

import (
	"os"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/pkg/errors"
)

// Transformation describes a generic operation that is meant
// to mutate a single cast at a time.
type Transformation interface {
	// Transform performs a mutation (in-place) in a cast.
	// note.: no concurrency mechanisms guarantee that this
	// is safe to be used across multiple goroutines
	Transform(c *cast.Cast) (err error)
}

// Transformer wraps the agents in a tranformation pipeline.
// Once created (see `New`), whenever a transformation is meant
// to be performed (see `Transform`), `Transformer` will read a
// complete cast from `intput`, process it with the trasnformation
// set and then output the resulting cast to `output`.
//
// We can illustrate the whole process like this:
//
//   input ==> transformation ==> output
//
// Note.: `input` will be consumed until EOF before the transformation
// is applied.
type Transformer struct {
	input          *os.File
	output         *os.File
	transformation Transformation
}

// New instantiates a new Transformer instance.
//
// The parameters are:
// - t: a Transformation interface implementor;
// - input: name of a file to read a cast from; and
// - output: name of a file to save the transformed cast to.
func New(t Transformation, input, output string) (m *Transformer, err error) {
	if t == nil {
		err = errors.Errorf("a transformation must be specified")
		return
	}

	m = &Transformer{
		input:          os.Stdin,
		output:         os.Stdout,
		transformation: t,
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

// Transform performs the central piece of the cast transformation process:
// 1. decodes a cast from `input`; then
// 2. applies the transformation in the cast that now lives in memory; then
// 3. encodes the cast, saving it to `output`.
func (m *Transformer) Transform() (err error) {
	var decodedCast *cast.Cast

	decodedCast, err = cast.Decode(m.input)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to decode cast from input")
		return
	}

	err = m.transformation.Transform(decodedCast)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to transform cast")
		return
	}

	err = cast.Encode(m.output, decodedCast)
	if err != nil {
		err = errors.Wrapf(err,
			"failed to save modified cast")
		return
	}

	return
}

// Close closes any open resources (input and output).
func (m *Transformer) Close() (err error) {
	if m.output != nil && m.output != os.Stdout {
		m.output.Close()
	}

	if m.input != nil && m.input != os.Stdin {
		m.input.Close()
	}

	return
}
