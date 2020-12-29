package reflections

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestReflections(t *testing.T) {

	t.Run("reflections tests", func(t *testing.T) {
		want := "Hello, Wolle" // expected
		var got []string

		x := struct {
			Name string
		}{want}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if got[0] != want {
			t.Errorf("got %q, want %q", got[0], want)
		}
		//fmt.Printf("Got is: %q", got)
		if len(got) != 1 {
			t.Errorf("wromg number of function calls, got %d want %d", len(got), 1)
		}
	})
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"walk through: Struct with one string field",
			struct {
				Name string
			}{"Wolle"},
			[]string{"Wolle"},
		},
		{
			"walk through: Struct with zwo string field",
			struct {
				Name string
				City string
			}{"Wolle", "Hamburg"},
			[]string{"Wolle", "Hamburg"},
		},
		{
			"walk through: Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Wolle", 56},
			[]string{"Wolle"},
		},
		{
			"walk through: Struct with nested field",
			Person{
				"Wolle",
				Profile{56, "Hamburg"},
			},
			[]string{"Wolle", "Hamburg"},
		},
		{
			"walk through: Pointer to Things",
			&Person{
				"Wolle",
				Profile{56, "Hamburg"},
			},
			[]string{"Wolle", "Hamburg"},
		},
		{
			"walk through: Slices",
			[]Profile{
				{33, "Berlin"},
				{35, "München"},
			},
			[]string{"Berlin", "München"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

}
