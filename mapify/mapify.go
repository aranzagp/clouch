package mapify

import (
	"errors"
	"reflect"
)

var (
	ErrNotAStruct = errors.New("Unsupported. Not a struct")
	clouch        = "clouch"
)

func Do(v interface{}) (map[string]interface{}, error) {
	typ := reflect.TypeOf(v).Elem()

	if !isStruct(typ) {
		return nil, ErrNotAStruct
	}
	// TODO: Validate we receive a pointer

	value := reflect.ValueOf(v).Elem()
	res, err := getStructFields(value, true)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func getStructFields(value reflect.Value, level bool) (map[string]interface{}, error) {
	if value.Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	res := map[string]interface{}{}
	typ := value.Type()

	idNum, err := idTagExists(typ)
	if err != nil {
		return nil, err
	}

	revNum, err := revTagExists(typ)
	if err != nil {
		return nil, err
	}

	for i := 0; i < value.NumField(); i++ {

		if idNum == i && level {
			continue
		}

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

		if revNum == i && level {
			if val.Kind() != reflect.Slice {
				return nil, errors.New("Revs is not an array")
			}

			rev := val.Index(0)
			res["_rev"] = rev.Interface()
			continue
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

			r, err := getStructFields(val, false)
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
