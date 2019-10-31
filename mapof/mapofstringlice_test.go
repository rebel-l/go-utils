package mapof_test

import (
	"testing"

	"github.com/rebel-l/go-utils/slice"

	"github.com/rebel-l/go-utils/mapof"
)

func TestStringSliceMap_GetValuesForKey(t *testing.T) {
	tests := []struct {
		name  string
		given mapof.StringSliceMap
		key   string
		want  slice.StringSlice
	}{
		{
			name:  "empty map",
			given: mapof.StringSliceMap{},
			key:   "first",
			want:  slice.StringSlice{},
		},
		{
			name:  "key not matching",
			given: mapof.StringSliceMap{"first": slice.StringSlice{"one"}},
			key:   "something",
			want:  slice.StringSlice{},
		},
		{
			name:  "key matches",
			given: mapof.StringSliceMap{"first": slice.StringSlice{"one"}},
			key:   "first",
			want:  slice.StringSlice{"one"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.given.GetValuesForKey(test.key)
			if len(test.want) != len(actual) {
				t.Errorf("expected %v but got %v", test.want, actual)
			}
		})
	}
}

func getKeyExistTestCases() []struct {
	name  string
	given mapof.StringSliceMap
	key   string
	want  bool
} {
	return []struct {
		name  string
		given mapof.StringSliceMap
		key   string
		want  bool
	}{
		{
			name:  "empty map",
			given: mapof.StringSliceMap{},
			key:   "first",
			want:  false,
		},
		{
			name:  "key not matching",
			given: mapof.StringSliceMap{"first": slice.StringSlice{"one"}},
			key:   "something",
			want:  false,
		},
		{
			name:  "key matches",
			given: mapof.StringSliceMap{"first": slice.StringSlice{"one"}},
			key:   "first",
			want:  true,
		},
	}
}

func TestStringSliceMap_KeyExists(t *testing.T) {
	tests := getKeyExistTestCases()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.given.KeyExists(test.key)
			if test.want != actual {
				t.Errorf("expected %t but got %t", test.want, actual)
			}
		})
	}
}

func TestStringSliceMap_KeyNotExists(t *testing.T) {
	tests := getKeyExistTestCases()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.given.KeyNotExists(test.key)
			if !test.want != actual {
				t.Errorf("expected %t but got %t", test.want, actual)
			}
		})
	}
}

func TestStringSliceMap_AddUniqueValue(t *testing.T) {
	tests := []struct {
		name  string
		given mapof.StringSliceMap
		key   string
		value string
		want  mapof.StringSliceMap
	}{
		{
			name:  "key no match",
			given: mapof.StringSliceMap{},
			key:   "first",
			value: "one",
			want:  mapof.StringSliceMap{"first": slice.StringSlice{"one"}},
		},
		{
			name:  "key match - no duplicate",
			given: mapof.StringSliceMap{"first": slice.StringSlice{"one"}},
			key:   "first",
			value: "two",
			want:  mapof.StringSliceMap{"first": slice.StringSlice{"one", "two"}},
		},
		{
			name:  "key match - duplicate",
			given: mapof.StringSliceMap{"first": slice.StringSlice{"one"}},
			key:   "first",
			value: "one",
			want:  mapof.StringSliceMap{"first": slice.StringSlice{"one"}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.given.AddUniqueValue(test.key, test.value)
			if len(test.want) != len(test.given) {
				t.Errorf("expected %v but got %v", test.want, test.given)
			}

			for key, want := range test.want {
				if test.given.KeyNotExists(key) {
					t.Errorf("expected %v but got %v", test.want, test.given)
				}

				a := test.given.GetValuesForKey(key)
				for _, w := range want {
					if a.IsNotIn(w) {
						t.Errorf("expected %v but got %v", test.want, test.given)
					}
				}

				if hasDuplicates(a) {
					t.Errorf("expected %v but got %v", test.want, test.given)
				}
			}
		})
	}
}

func hasDuplicates(s slice.StringSlice) bool {
	counter := make(map[string]bool)
	for _, v := range s { // nolint: wsl
		if _, ok := counter[v]; ok {
			return true
		}

		counter[v] = true
	}

	return false
}
