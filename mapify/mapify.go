package mapify

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNotAStruct = errors.New("Unsupported. Not a struct")
	clouch        = "clouch"
)

func Do(v interface{}) (map[string]interface{}, error) {
	typ := reflect.TypeOf(v).Elem()

	if typ.Kind() != reflect.Struct {
		fmt.Println(typ.Kind())
		return nil, ErrNotAStruct

	}

	// TODO: Validate we receive a pointer

	value := reflect.ValueOf(v).Elem()
	res, err := getStructFields(value)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getStructFields(value reflect.Value) (map[string]interface{}, error) {
	if value.Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	res := map[string]interface{}{}
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {

		val := value.Field(i)
		tp := typ.Field(i)

		name := ""
		tagField := tp.Tag.Get(clouch)

		tg := getTag(tagField)

		if tagField != tg.name {
			name = tagField
		} else {
			name = tp.Name
		}

		if tg.Ignore() {
			continue
		}

		if tg.name == "_id" {
			continue
		}

		if tg.name == "_revs" {
			name = "_rev"
		}

		if tp.Type.Kind() == reflect.Ptr {
			if p := val.Pointer(); p == 0 {

				if tg.OmitEmpty() {
					continue
				}

				res[name] = nil
				continue
			}

			val = val.Elem()
		}

		if val.Kind() == reflect.Struct {

			r, err := getStructFields(val)
			if err != nil {
				return nil, err
			}

			res[name] = r
			continue
		}

		res[name] = val.Interface()

	}

	return res, nil
}
