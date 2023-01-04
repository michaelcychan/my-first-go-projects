package reflection

import (
	"reflect"
)

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

	var numOfField int = 0

	// this is to declare a variable,
	// this variable will be a function taking an int and return a reflect.Value
	var getField func(int) reflect.Value

	switch val.Kind() {
	// if x is already string, just run the function
	case reflect.String:
		fn(val.String())

	// // .NumField() and .Field() will cause panic if not used on struct
	// // that is why in our test case, the simplest example is still a struct
	// case reflect.Struct:
	// 	for i := 0; i < val.NumField(); i++ {
	// 		field := val.Field(i)
	// 		walk(field.Interface(), fn)
	// 	}

	// // Slice does not take .NumField() or .Field(), so we need another case
	// // because val is a reflect object, we need to use .Len() and .Index(i)
	// case reflect.Slice:
	// 	for i := 0; i < val.Len(); i++ {
	// 		field := val.Index(i)
	// 		walk(field.Interface(), fn)
	// 	}
	// 	return

	// .NumField() and .Field() will cause panic if not used on struct
	// that is why in our test case, the simplest example is still a struct

	// instead of the old method,
	// the refactored code keeps only getting values and fields num in switch
	case reflect.Struct:
		numOfField = val.NumField()
		getField = val.Field

	// Slice does not take .NumField() or .Field(), so we need another case
	// because val is a reflect object, we need to use .Len() and .Index(i)
	case reflect.Slice, reflect.Array:
		numOfField = val.Len()
		getField = val.Index
	case reflect.Map:

		// the following code will print keys rather than values
		// for _, key := range val.MapKeys() {
		// 	fn(key.String())
		// }

		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	// this is outside switch
	// reflect.String will have numOfField == 0, so no worry
	// for slice and struct, they get the numOfField inside switch,
	// and they their respective .Field() and .Index() is copied to getField

	for i := 0; i < numOfField; i++ {
		walk(getField(i).Interface(), fn)
	}

}
