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

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	// obsolete
	// // this is to declare a variable,
	// // this variable will be a function taking an int and return a reflect.Value
	// var getField func(int) reflect.Value
	// var numOfField int = 0

	switch val.Kind() {
	// if x is already string, just run the function
	case reflect.String:
		fn(val.String())

	// .NumField() and .Field() will cause panic if not used on struct
	// that is why in our test case, the simplest example is still a struct
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}

	// Slice does not take .NumField() or .Field(),
	// because val is a reflect object, we need to use .Len() and .Index(i)
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
		return

	// // the below was intermin refactor,
	// // the refactored code keeps only getting values and fields num in switch
	// // but this way does not work for map, so reverted back
	// case reflect.Struct:
	// 	numOfField = val.NumField()
	// 	getField = val.Field

	// // Slice does not take .NumField() or .Field(), so we need another case
	// // because val is a reflect object, we need to use .Len() and .Index(i)
	// case reflect.Slice, reflect.Array:
	// 	numOfField = val.Len()
	// 	getField = val.Index
	case reflect.Map:
		// the following code will print keys rather than values
		// for _, key := range val.MapKeys() {
		// 	fn(key.String())
		// }
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	case reflect.Chan:
		// https://go.dev/ref/spec#Receive_operator
		// second return value (ok) is a bool
		// The value of ok is true if the value received was delivered by a successful send operation to the channel,
		// or false if it is a zero value generated because the channel is closed and empty.
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		// cannot directly call function, but using .Call()
		valReturn := val.Call(nil)
		for _, res := range valReturn {
			walkValue(res)
		}

	}

	// this is outside switch
	// reflect.String will have numOfField == 0, so no worry
	// for slice and struct, they get the numOfField inside switch,
	// and they their respective .Field() and .Index() is copied to getField

	// obsolete
	// for i := 0; i < numOfField; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }

}
