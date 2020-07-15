package utils

import (
	"reflect"
	"strings"
)

// StructToMap ...
func StructToMap(item interface{}) map[string]interface{} {
	results := map[string]interface{}{}
	if item == nil {
		return results
	}
	v := reflect.TypeOf(item)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get("json")
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" && field != reflect.Zero(reflect.TypeOf(field)).Interface() {
			tag = strings.Replace(tag, ",omitempty", "", -1)
			if v.Field(i).Type.Kind() == reflect.Struct {
				results[tag] = StructToMap(field)
			} else {
				results[tag] = field
			}
		}
	}

	return results
}
