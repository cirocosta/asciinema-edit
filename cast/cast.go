package cast

import (
	_ "encoding/json"
)

type Header struct {
	Version   uint8  `json:"version,omitempty"`
	Width     uint   `json:"width,omitempty"`
	Height    uint   `json:"height,omitempty"`
	Timestamp uint   `json:"timestamp,omitempty"`
	Title     string `json:"title,omitempty"`
	Env       struct {
		Shell string `json:"SHELL,omitempty"`
		Term  string `json:"TERM,omitempty"`
	} `json:"env,omitempty"`
}

type Event struct {
	Time float64
	Type string
	Data string
}

type Cast struct {
	Header      Header
	EventStream []*Event
}
