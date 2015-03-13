package mapify

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrNotAStruct = errors.New("Unsupported. Not a struct")
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

		if tp.Type.Kind() == reflect.Ptr {
			if p := val.Pointer(); p == 0 {
				res[tp.Name] = nil
				continue
			}

			val = val.Elem()
		}

		if val.Kind() == reflect.Struct {

			r, err := getStructFields(val)
			if err != nil {
				return nil, err
			}

			res[tp.Name] = r
			continue
		}

		res[tp.Name] = val.Interface()

	}

	return res, nil
}
