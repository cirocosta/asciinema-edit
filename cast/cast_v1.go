package cast

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

type CastV1 struct {
	Header
	Duration    float64
	EventStream []*Event
}

type castV1Json struct {
	Header
	Duration float64         `json:"duration"`
	Stdout   [][]interface{} `json:"stdout"`
}

func DecodeV1(reader io.Reader) (cast *CastV1, err error) {
	if reader == nil {
		err = errors.Errorf("reader must not be nil")
		return
	}

	var (
		jsonObj = new(castV1Json)
		decoder *json.Decoder
	)

	decoder = json.NewDecoder(reader)

	err = decoder.Decode(&jsonObj)
	if err != nil {
		err = errors.Wrapf(err,
			"couldn't decode cast v1")
		return
	}

	cast = &CastV1{
		Header: jsonObj.Header,
		Duration: jsonObj.Duration,
		EventStream: make([]Event, len(jsonObj.Stdout)),
	}

	var (
		delta float64
		data string
		ok bool
	)
	for idx, out := range jsonObj.Stdout {
		delta, ok = out[0].(float64)
		if !ok {
			err = errors.Errorf("first element of event is not a float64 (delta)")
			return
		}

		data, ok = out[1].(string)
		if !ok {
			err = errors.Errorf("second element of event is not a string (data)")
			return
		}

		// TODO fix this
		cast.EventStream[idx] = Event{
			Type: "o",
			Time: delta,
		}

	}


	return
}
