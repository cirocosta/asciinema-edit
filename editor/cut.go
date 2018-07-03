package editor

import (
	"math"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/pkg/errors"
)

// Cut removes a piece of the cast event stream as specified by
// `from` and `to`.
//
// It assumes that the provided `cast` is entirely valid (see
// `github.com/cirocosta/asciinema-edit/cast#Validate`).
//
// If `from == to`:
//	the exact timestamp is removed.
//
// If `from < to`:
//	all timestamps from `from` to `to` are removed (both included).
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

	var (
		fromIdx = -1
		toIdx   = -1
	)

	for idx, ev := range c.EventStream {
		if ev.Time == from {
			fromIdx = idx
		}

		if ev.Time == to {
			toIdx = idx
		}
	}

	if fromIdx == -1 {
		err = errors.Errorf("couldn't find initial frame")
		return
	}

	if toIdx == -1 {
		err = errors.Errorf("couldn't find final frame")
		return
	}

	if toIdx+1 < len(c.EventStream) {
		delta := c.EventStream[toIdx+1].Time - c.EventStream[fromIdx].Time
		for _, remainingElem := range c.EventStream[toIdx+1:] {
			remainingElem.Time -= delta
			remainingElem.Time = math.Round(remainingElem.Time*1000) / 1000
		}
	}

	c.EventStream = append(
		c.EventStream[:fromIdx],
		c.EventStream[toIdx+1:]...)

	return
}
