package mapify

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	NotAStruct = errors.New("Unsupported. Not a struct")
)

func Do(v interface{}) (map[string]interface{}, error) {
	values := reflect.TypeOf(v).Elem()

	if values.Kind() != reflect.Struct {
		fmt.Println(values.Kind())
		return nil, NotAStruct

	}

	res := map[string]interface{}{}
	for i := 0; i < values.NumField(); i++ {
		val := values.Field(i)
		res[val.Name] = reflect.ValueOf(v).Elem().FieldByName(val.Name).Interface()
	}
	return res, nil
}
