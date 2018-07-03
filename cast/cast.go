// Package cast contains the essential structures for dealing with
// an asciinema cast of the v2 format.
//
// The current implementation is based on the V2 format as of July 2nd, 2018.
//
// From [1], asciicast v2 file is a newline-delimited JSON file where:
//
// - first line contains header (initial terminal size, timestamp and other
//   meta-data), encoded as JSON object; and
// - all following lines form an event stream, each line representing a separate
//   event, encoded as 3-element JSON array.
//
// [1]: https://github.com/asciinema/asciinema/blob/49a892d9e6f57ab3a774c0835fa563c77cf6a7a7/doc/asciicast-v2.md.
package cast

import (
	"io"

	"github.com/pkg/errors"
)

// Header represents the asciicast header - a JSON-encoded object containing
// recording meta-data.
type Header struct {
	// Version represents the version of the current ascii cast format
	// (must be `2`).
	//
	// This field is required for a valid header.
	Version uint8 `json:"version"`

	// With is the initial terminal width (number of columns).
	//
	// This field is required for a valid header.
	Width uint `json:"width"`

	// Height is the initial terminal height (number of rows).
	//
	// This field is required for a valid header.
	Height uint `json:"height"`

	// Timestamp is the unix timestamp of the beginning of the
	// recording session.
	Timestamp uint `json:"timestamp,omitempty"`

	// Command corresponds to the name of the command that was
	// recorded.
	Command string `json:"command,omitempty"`

	// Theme describes the color theme of the recorded terminal.
	Theme struct {
		// Fg corresponds to the normal text color (foreground).
		Fg string `json:"fg,omitempty"`

		// Bg corresponds to the normal background color.
		Bg string `json:"bg,omitempty"`

		// Palette specifies a list of 8 or 16 colors separated by
		// colon character to apply a theme to the session
		Palette string `json:"palette,omitempty"`
	} `json:"theme,omitempty"`

	// Title corresponds to the title of the cast.
	Title string `json:"title,omitempty"`

	// IdleTimeLimit specifies the maximum amount of idleness between
	// one command and another.
	IdleTimeLimit float64 `json:"title,omitempty"`

	// Env specifies a map of environment variables captured by the
	// asciinema command.
	//
	// ps.: the official asciinema client only captures `SHELL` and `TERM`.
	Env struct {
		// Shell corresponds to the captured SHELL environment variable.
		Shell string `json:"SHELL,omitempty"`

		// Term corresponds to the captured TERM environment variable.
		Term string `json:"TERM,omitempty"`
	} `json:"env,omitempty"`
}

// Event represents terminal inputs that get recorded by asciinema.
type Event struct {
	// Time indicates when this event happened, represented as the number
	// of seconds since the beginning of the recording session.
	Time float64

	// Type represents the type of the data that's been recorded.
	//
	// Two types are possible:
	// - "o": data written to stdout; and
	// - "i": data read from stdin.
	Type string

	// Data represents the data recorded from the terminal.
	Data string
}

// Cast represents the whole asciinema session.
type Cast struct {
	// Header presents the recording metadata.
	Header Header

	// EventStream contains all the events that were generated during
	// the recording.
	EventStream []Event
}

// Decode reads the whole contents of the reader passed as argument, validates
// whether the stream contains a valid asciinema cast and then unmarshals it
// into a cast struct.
func Decode(r io.Reader) (cast *Cast, err error) {
	return
}

// ValidateHeader verifies whether the provided `cast` header structure is valid
// or not based on the asciinema cast v2 protocol.
func ValidateHeader(header *Header) (isValid bool, err error) {
	if header.Version != 2 {
		err = errors.Errorf("only casts with version 2 are valid")
		return
	}

	if header.Width == 0 {
		err = errors.Errorf("a valid width (>0) must be specified")
		return
	}

	if header.Height == 0 {
		err = errors.Errorf("a valid height (>0) must be specified")
		return
	}

	isValid = true
	return
}

// ValidateEvent checks whether the provided `Event` is properly formed.
func ValidateEvent(event *Event) (isValid bool, err error) {
	return
}

// Encode writes the encoding of `Cast` into the writer passed as an argument.
func Encode(writer io.Writer, cast *Cast) (err error) {
	if writer == nil {
		err = errors.Errorf("a writer must be specified")
		return
	}

	if cast == nil {
		err = errors.Errorf("a cast must be specified")
		return
	}

	return
}
