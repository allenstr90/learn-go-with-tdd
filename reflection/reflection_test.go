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
			struct{ Name string }{"Bruce"},
			[]string{"Bruce"},
		},
		{
			"struct with 2 string fields",
			struct {
				Name     string
				LastName string
			}{"Bruce", "Wayne"},
			[]string{"Bruce", "Wayne"},
		},
		{
			"struct with one no string field",
			struct {
				Name string
				Age  int
			}{"Bruce", 33},
			[]string{"Bruce"},
		},
		{
			"nested fields",
			Person{"Bruce", Info{33, "Gotham"}},
			[]string{"Bruce", "Gotham"},
		},
		{
			"using pointers",
			&Person{"Bruce", Info{33, "Gotham"}},
			[]string{"Bruce", "Gotham"},
		},
		{
			"slices",
			[]Info{
				{33, "Gotham"},
				{34, "New York"},
			},
			[]string{"Gotham", "New York"},
		},
		{
			"arrays",
			[2]Info{
				{33, "Gotham"},
				{34, "New York"},
			},
			[]string{"Gotham", "New York"},
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
	t.Run("maps due order issue", func(t *testing.T) {
		aMap := map[int]string{
			33: "Gotham",
			34: "New York",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Gotham")
		assertContains(t, got, "New York")
	})

	t.Run("channels", func(t *testing.T) {
		aChannel := make(chan Info)

		go func() {
			aChannel <- Info{33, "Gotham"}
			aChannel <- Info{34, "New York"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Gotham", "New York"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("functions", func(t *testing.T) {
		aFunction := func() (Info, Info) {
			return Info{33, "Gotham"}, Info{34, "New York"}
		}

		var got []string
		want := []string{"Gotham", "New York"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
