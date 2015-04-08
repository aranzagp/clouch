package unmapify

import (
	"fmt"
	"reflect"

	"github.com/thetonymaster/clouch/utils"
)

func Do(v interface{}, mp map[string]interface{}) error {
	typ := reflect.TypeOf(v).Elem()

	if !utils.IsStruct(typ) {
		return utils.ErrNotAStruct
	}

	value := reflect.ValueOf(v).Elem()
	err := setStructFields(&value, mp)
	if err != nil {
		return err
	}
	return nil
}

func setStructFields(value *reflect.Value, mp map[string]interface{}) error {

	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {

		tagStr := typ.Field(i).Tag.Get("clouch")
		tg := utils.GetTag(tagStr)

		name := ""
		if tg.Name != "" {
			name = tg.Name
		} else {
			name = typ.Field(i).Name
		}

		if vl, ok := mp[name]; ok {
			//do something here
			fmt.Print("Name" + name)
			val := reflect.ValueOf(vl)
			value.Field(i).Set(val)
		}

	}

	return nil
}
