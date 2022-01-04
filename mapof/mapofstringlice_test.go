package mapof_test

import (
	"testing"

	"github.com/rebel-l/go-utils/mapof"
	"github.com/rebel-l/go-utils/slice"
)

func TestStringSliceMap_GetValuesForKey(t *testing.T) {
	t.Parallel()

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
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.given.GetValuesForKey(tc.key)
			if len(tc.want) != len(actual) {
				t.Errorf("expected %v but got %v", tc.want, actual)
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
	t.Parallel()

	tests := getKeyExistTestCases()

	for _, test := range tests {
		tc := test
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.given.KeyExists(tc.key)
			if tc.want != actual {
				t.Errorf("expected %t but got %t", tc.want, actual)
			}
		})
	}
}

func TestStringSliceMap_KeyNotExists(t *testing.T) {
	t.Parallel()

	tests := getKeyExistTestCases()

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.given.KeyNotExists(tc.key)
			if !tc.want != actual {
				t.Errorf("expected %t but got %t", tc.want, actual)
			}
		})
	}
}

func TestStringSliceMap_AddUniqueValue(t *testing.T) {
	t.Parallel()

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
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			tc.given.AddUniqueValue(tc.key, tc.value)
			if len(tc.want) != len(tc.given) {
				t.Errorf("expected %v but got %v", tc.want, tc.given)
			}

			for key, want := range tc.want {
				if tc.given.KeyNotExists(key) {
					t.Errorf("expected %v but got %v", tc.want, tc.given)
				}

				a := tc.given.GetValuesForKey(key)
				for _, w := range want {
					if a.IsNotIn(w) {
						t.Errorf("expected %v but got %v", tc.want, tc.given)
					}
				}

				if hasDuplicates(a) {
					t.Errorf("expected %v but got %v", tc.want, tc.given)
				}
			}
		})
	}
}

func hasDuplicates(s slice.StringSlice) bool {
	counter := make(map[string]bool)
	for _, v := range s {
		if _, ok := counter[v]; ok {
			return true
		}

		counter[v] = true
	}

	return false
}
