// From Twitter
// golang challenge:
// write a function walk(x interface{}, fn func(string))
// which takes a struct x and
// calls fn for all strings fields found inside.
// difficulty level: recursively.

package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	type Profile struct {
		Age  int
		City string
	}
	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		// {
		// 	"struct with one string field (intentionally wrong)",
		// 	struct {
		// 		Name string
		// 	}{"Wrong Input"}, // actually used as parameter in the spy function
		// 	[]string{"Output"}, // expected value used in spy function
		// },
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non-string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		// {
		// 	"old nested fields without pre-defined type",
		// 	struct {
		// 		Name string
		// 		Profile struct {
		// 			Age int
		// 			City string
		// 		}
		// 	}{
		// 		"Chris",
		// 		struct{
		// 			Age int
		// 			City string
		// 		}{33, "London"},
		// 	},
		// 	[]string{"Chirs", "London"},
		// },
		{
			"struct with nested fields",
			Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			// func(input string) {...} is an anonymous function put into walk func as an argument
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("expected input processed: %v, but got %v", test.ExpectedCalls, got)
			}

		})
	}
}
