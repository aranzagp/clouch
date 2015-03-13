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

	res := map[string]interface{}{}
	vValues := reflect.ValueOf(v).Elem()

	for i := 0; i < typ.NumField(); i++ {
		tp := typ.Field(i)

		if tp.Type.Kind() == reflect.Struct {

			r, err := getStructFields(vValues.Field(i))
			if err != nil {
				return nil, err
			}

			res[tp.Name] = r
			continue
		}

		if tp.Type.Kind() == reflect.Ptr {
			if p := vValues.FieldByName(tp.Name).Pointer(); p == 0 {
				res[tp.Name] = nil
			} else {
				res[tp.Name] = vValues.FieldByName(tp.Name).Elem().Interface()

			}
		} else {
			res[tp.Name] = vValues.FieldByName(tp.Name).Interface()
		}

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
		res[tp.Name] = val.Interface()

	}

	return res, nil
}
