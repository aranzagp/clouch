package utils

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNoID       = errors.New("No ID tag Found")
	ErrNoRevs     = errors.New("No Rev tag Found")
	ErrNotAStruct = errors.New("Unsupported. Not a struct")
)

const (
	Ignore    = "-"
	OmitEmpty = "omitempty"
)

func IsStruct(typ reflect.Type) bool {

	if typ.Kind() != reflect.Struct {
		fmt.Println(typ.Kind())
		return false

	}
	return true
}
