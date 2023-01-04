package reflection

import "reflect"

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// handling x if it is a pointer
	// .Elem() can get the value of which the pointer points to
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {

	// if x is already string, just run the function
	case reflect.String:
		fn(val.String())

	// .NumField() and .Field() will cause panic if not used on struct
	// that is why in our test case, the simplest example is still a struct
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			walk(field.Interface(), fn)
		}

	// Slice does not take .NumField() or .Field(), so we need another case
	// because val is a reflect object, we need to use .Len() and .Index(i)
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			field := val.Index(i)
			walk(field.Interface(), fn)
		}
		return

	}

}
