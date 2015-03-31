package utils

import "errors"

var (
	ErrNoID   = errors.New("No ID tag Found")
	ErrNoRevs = errors.New("No Rev tag Found")
)

const (
	Ignore    = "-"
	OmitEmpty = "omitempty"
)
