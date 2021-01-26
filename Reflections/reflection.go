package Reflections

import "reflect"

func Walk(x interface{},fn func(string2 string)) {

	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		for i:=0;i<val.NumField();i++ {
			Walk(val.Field(i).Interface(),fn)
		}
	case reflect.Slice,reflect.Array:
		for i:=0; i< val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walk(val.MapIndex(key).Interface(), fn)
		}
	}
}
