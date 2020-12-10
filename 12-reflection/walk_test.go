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
			"Struct with one string field",
			struct {
				Name string
			}{"LaughingMan"},
			[]string{"LaughingMan"},
		}, {
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"LaughingMan", "Tokyo"},
			[]string{"LaughingMan", "Tokyo"},
		}, {
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"LaughingMan", 33},
			[]string{"LaughingMan"},
		}, {
			"Nested fields",
			Person{"LaughingMan",
				Profile{33, "Tokyo"},
			},
			[]string{"LaughingMan", "Tokyo"},
		}, {
			"Pointers to things",
			&Person{
				"LaughingMan",
				Profile{33, "Tokyo"},
			},
			[]string{"LaughingMan", "Tokyo"},
		}, {
			"Slices",
			[]Profile{
				{33, "Tokyo"},
				{34, "Manchester"},
			},
			[]string{"Tokyo", "Manchester"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var actual []string
			walk(test.Input, func(input string) {
				actual = append(actual, input)
			})

			if !reflect.DeepEqual(actual, test.ExpectedCalls) {
				t.Errorf("actual %v, expected %v", actual, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
