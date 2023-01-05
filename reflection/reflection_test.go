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
		{
			"using pointers",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"handling slices",
			[]Profile{
				{33, "London"},
				{40, "Hong Kong"},
			},
			[]string{"London", "Hong Kong"},
		}, {
			"handling array",
			[3]Profile{
				{33, "London"},
				{40, "Hong Kong"},
				{48, "Tokyo"},
			},
			[]string{"London", "Hong Kong", "Tokyo"},
			// }, {
			//
			// this way of testing map might fail because of sequence
			// so testing map is done outside this struct test
			//
			// 	"handling maps",
			// 	map[string]string{
			// 		"City":      "Hong Kong",
			// 		"Continent": "Asia",
			// 	},
			// 	[]string{"Hong Kong", "Asia"},
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
	t.Run("testing with map", func(t *testing.T) {
		var got []string
		inputMap := map[string]string{
			"City":      "Hong Kong",
			"Continent": "Asia",
		}
		walk(inputMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Hong Kong")
		assertContains(t, got, "Asia")

	})
	t.Run("testing with channel", func(t *testing.T) {
		testChannel := make(chan Profile)

		go func() {
			testChannel <- Profile{30, "London"}
			testChannel <- Profile{35, "Hong Kong"}
			close(testChannel)
		}()

		var got []string

		walk(testChannel, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Hong Kong")
		assertContains(t, got, "London")
	})
	t.Run("testing with func", func(t *testing.T) {
		// a mock function returns two Profile
		testFunction := func() (Profile, Profile) {
			return Profile{33, "Tokyo"}, Profile{35, "Osaka"}
		}
		var got []string
		expected := []string{"Tokyo", "Osaka"}

		walk(testFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected %v, but got %v", expected, got)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, v := range haystack {
		if v == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %s, but does not", haystack, needle)
	}
}
