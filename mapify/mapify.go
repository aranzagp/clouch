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
	values := reflect.TypeOf(v).Elem()

	if values.Kind() != reflect.Struct {
		fmt.Println(values.Kind())
		return nil, ErrNotAStruct

	}

	res := map[string]interface{}{}
	vValues := reflect.ValueOf(v).Elem()

	for i := 0; i < values.NumField(); i++ {
		val := values.Field(i)

		if val.Type.Kind() == reflect.Ptr {
			if p := vValues.FieldByName(val.Name).Pointer(); p == 0 {
				res[val.Name] = nil
			} else {
				res[val.Name] = vValues.FieldByName(val.Name).Elem().Interface()

			}
		} else {
			res[val.Name] = vValues.FieldByName(val.Name).Interface()
		}

	}
	return res, nil
}
