package editor

import (
	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/pkg/errors"
)

// Cut removes a piece of the cast event stream as specified by
// `from` and `to`.
//
// It assumes that the provided `cast` is entirely valid (see
// `github.com/cirocosta/asciinema-edit/cast#Validate`).
//
// 1. search a time that is close to `from`; then
// 2. search a time that is close to `to`; then
// 3. remove all in between; then
// 4. adjust the time of the remaining.
func Cut(c *cast.Cast, from, to float64) (err error) {
	if c == nil {
		err = errors.Errorf("a cast must be specified")
		return
	}

	if len(c.EventStream) == 0 {
		err = errors.Errorf(
			"a cast with non-empty event stream must be supplied")
		return
	}

	if from > to {
		err = errors.Errorf(
			"`from` cant be bigger than `to`")
		return
	}

	return
}
