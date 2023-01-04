package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	// handling x as a pointer
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}

		// old if clauses replaced by switch cases
		// if field.Kind() == reflect.String {
		// 	fn(field.String())
		// }
		// if field.Kind() == reflect.Struct {
		// 	walk(field.Interface(), fn)
		// }
	}

}
