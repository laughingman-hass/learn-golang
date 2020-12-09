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
        },{
            "Struct with two string fields",
            struct {
                Name string
                City string
            }{"LaughingMan", "Tokyo"},
            []string{"LaughingMan", "Tokyo"},
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
