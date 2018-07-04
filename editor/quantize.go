package editor

import (
	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/pkg/errors"
)

// QuantizeRange describes a quantization range.
//
// When applied to a quantization function, any values that lie in the
// range are trimmed to `from`.
type QuantizeRange struct {
	// From indicates the start of the range
	From float64
	// To indicates the end of the range
	To float64
}

// NewQuantizeRange creates a new quantize range performing some basic
// validations:
// - `from` < `to`
func NewQuantizeRange(from, to float64) (q *QuantizeRange, err error) {
	if from < to {
		err = errors.Errorf("constraint not satisfied: from < to")
		return
	}

	q = &QuantizeRange{
		From: from,
		To:   to,
	}
	return
}

// InRange verifies whether a given value lies within the quantization
// range.
func (q *QuantizeRange) InRange(value float64) bool {
	return value >= q.From && value < q.To
}

// Overlaps verifies whether a given range (`another`) overlaps with
// this range.
func (q *QuantizeRange) RangeOverlaps(another QuantizeRange) bool {
	return q.InRange(another.From) || q.InRange(another.To)
}

// Quantize constraints a set of inputs that lie in a range to a single
// value that corresponds to the lower bound of such range.
//
// For instance, consider the following timestamps:
//
//	 1  2  5  9 10 11
//
// Assuming that we quantize over [2,6), we'd cut any delays between 2 and
// 6 seconds to 2 second:
//
//	 1  2  4  6  7  8
//
// This can be more easily visualized by looking at the delay quantization:
//
//      delta = 1.000000 | qdelta = 1.000000
//      delta = 3.000000 | qdelta = 2.000000
//      delta = 4.000000 | qdelta = 2.000000
//      delta = 1.000000 | qdelta = 1.000000
//      delta = 1.000000 | qdelta = 1.000000
//
// The euristic is:
//
// 1. capture all delays
// 2. for each delay, check if it's within an acceptable delay range
// 3. if it fits, reduce the delay to the maximum allowed (floor of
//    the quantization range).
// 4. adjust the rest of the event stream.
func Quantize(c *cast.Cast, ranges []QuantizeRange) (err error) {
	if c == nil {
		err = errors.Errorf("cast must not be nil")
		return
	}

	if len(c.EventStream) == 0 {
		err = errors.Errorf("event stream must not be empty")
		return
	}

	if len(ranges) == 0 {
		err = errors.Errorf("at least one quantization range must be specified")
		return
	}

	var (
		deltas = make([]float64, len(c.EventStream))
		delta  float64
		i      int
	)

	for i = 0; i < len(c.EventStream)-1; i++ {
		delta = c.EventStream[i+1].Time - c.EventStream[i].Time

		for _, qRange := range ranges {
			if !qRange.InRange(delta) {
				continue
			}

			delta = qRange.From
			break
		}

		deltas[i] = delta
	}

	for i = 0; i < len(c.EventStream)-1; i++ {
		c.EventStream[i+1].Time = c.EventStream[i].Time + deltas[i]
	}

	return
}
