package editor

import (
	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/pkg/errors"
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

	var (
		i                int
		k                int
		delta            float64
		newDelta         float64
		accumulatedDelta float64
		deltas           = make([]float64, toIdx-fromIdx)
	)

	k = 0
	for i = fromIdx; i < toIdx; i++ {
		delta = c.EventStream[i+1].Time - c.EventStream[i].Time
		newDelta = delta * factor
		accumulatedDelta += (newDelta - delta)

		deltas[k] = newDelta
		k += 1
	}

	k = 0
	for i = fromIdx; i < toIdx; i++ {
		c.EventStream[i+1].Time = c.EventStream[i].Time + deltas[k]
		k += 1
	}

	if toIdx+1 < len(c.EventStream) {
		for _, remainingElem := range c.EventStream[toIdx+1:] {
			remainingElem.Time += accumulatedDelta
		}
	}

	return
}
