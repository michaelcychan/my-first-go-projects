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
		{
			"struct with one string field (intentionally wrong)",
			struct {
				Name string
			}{"Wrong Input"},
			[]string{"Output"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("expected input processed: %v, but got %v", test.ExpectedCalls, got)
			}

		})
	}
}
