package mapify

import (
	"errors"
	"reflect"

	"github.com/thetonymaster/clouch/utils"
)

var (
	clouch = "clouch"
)

func Do(v interface{}) (map[string]interface{}, error) {
	typ := reflect.TypeOf(v).Elem()

	if !utils.IsStruct(typ) {
		return nil, utils.ErrNotAStruct
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
		return nil, utils.ErrNotAStruct
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
<<<<<<< HEAD
		tg := getTag(tagField)

		if tagField != ",omitempty" {
			if tagField != tg.name {
				name = tagField
			} else {
				name = tp.Name
			}
=======

		tg := utils.GetTag(tagField)

		if tagField != tg.Name {
			name = tagField
		} else {
			name = tp.Name
>>>>>>> upstream/master
		}

		if tg.Ignore() {
			continue
		}

		if tagField == ",omitempty" {
			name = tp.Name
			switch {

			case isSlice(i, value):
				if val.Len() != 0 {
					res[name] = val.Interface()
				}

			case isString(i, value):
				if val.String() != "" {
					res[name] = val.Interface()
				}

			case isInt(i, value):
				if val.Int() != 0 {
					res[name] = val.Interface()
				}

			case isFloat(i, value):
				if val.Float() != 0 {
					res[name] = val.Float()
				}

			case isBool(i, value):
				if val.Bool() {
					res[name] = val.Bool()
				}

			case isPtr(i, value):
				if p := val.Pointer(); p != 0 {
					res[name] = val.Interface()
				}

			case val.Kind() == reflect.Struct:
				r, err := getStructFields(val, false)
				if err != nil {
					return nil, err
				}
				res[name] = r
				//continue
			}

		} else {

			switch {
			case isPtr(i, value) || isString(i, value) || isSlice(i, value):
				res[name] = val.Interface()

			case isFloat(i, value):
				res[name] = val.Float()

			case isInt(i, value):
				res[name] = val.Int()

			case isBool(i, value):
				res[name] = val.Bool()
			}

		}

		if revNum == i && level {
			if val.Kind() != reflect.Slice {
				return nil, errors.New("Revs is not an array")
			}

			if val.Len() < 1 {
				continue
			}
			rev := val.Index(0)
			res["_rev"] = rev.Interface()
			continue
		}

		// if tp.Type.Kind() == reflect.Ptr {
		// 	if p := val.Pointer(); p == 0 {
		//
		// 		if tg.OmitEmpty() {
		// 			continue
		// 		}
		//
		// 		res[name] = nil
		// 		continue
		// 	}
		//
		// 	val = val.Elem()
		// }

		//res[name] = val.Interface()

	}

	return res, nil
}
