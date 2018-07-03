package editor

import (
	"github.com/pkg/errors"

	"github.com/cirocosta/asciinema-edit/cast"
)

// Speed updates the cast speed by multiplying all of the
// timestamps in a given range by a given factor.
func Speed(c *cast.Cast, factor, from, to float64) (err error) {
	if c == nil {
		err = errors.Errorf("cast must not be nil")
		return
	}

	if len(c.EventStream) == 0 {
		err = errors.Errorf("event stream must be nonempty")
		return
	}

	if factor > 10 || factor < 0.1 {
		err = errors.Errorf("factor must be within 0.1 and 10 range")
		return
	}

	if from >= to {
		err = errors.Errorf("`from` must not be greater or equal than `to`")
		return
	}

	return
}
