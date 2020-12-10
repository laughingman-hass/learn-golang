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
		}, {
			"Arrays",
			[2]Profile{
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"LaughingMan": "Tokyo",
			"Hass":        "Manchester",
		}

		var actual []string
		walk(aMap, func(input string) {
			actual = append(actual, input)
		})

		assertContains(t, actual, "Tokyo")
		assertContains(t, actual, "Manchester")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Tokyo"}
			aChannel <- Profile{34, "Manchester"}
			close(aChannel)
		}()

		var actual []string
		expected := []string{"Tokyo", "Manchester"}

		walk(aChannel, func(input string) {
			actual = append(actual, input)
		})

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual %v, expected %v", actual, expected)
		}
	})
}

func assertContains(t *testing.T, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
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
