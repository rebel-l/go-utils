package slice_test

import (
	"testing"

	"github.com/rebel-l/go-utils/slice"
)

func getIsInTestCases() []struct {
	name   string
	given  slice.StringSlice
	search string
	want   bool
} {
	return []struct {
		name   string
		given  slice.StringSlice
		search string
		want   bool
	}{
		{
			name:   "empty slice",
			given:  slice.StringSlice{},
			search: "something",
			want:   false,
		},
		{
			name:   "one element - match",
			given:  slice.StringSlice{"one"},
			search: "one",
			want:   true,
		},
		{
			name:   "one element - no match",
			given:  slice.StringSlice{"one"},
			search: "something",
			want:   false,
		},
		{
			name:   "two elements - first matches",
			given:  slice.StringSlice{"one", "two"},
			search: "one",
			want:   true,
		},
		{
			name:   "two elements - second matches",
			given:  slice.StringSlice{"one", "two"},
			search: "two",
			want:   true,
		},
		{
			name:   "two elements - no match",
			given:  slice.StringSlice{"one", "two"},
			search: "something",
			want:   false,
		},
		{
			name:   "empty search, empty slice",
			given:  slice.StringSlice{},
			search: "",
			want:   false,
		},
		{
			name:   "empty search, one element - no match",
			given:  slice.StringSlice{"something"},
			search: "",
			want:   false,
		},
		{
			name:   "empty search, one element - match",
			given:  slice.StringSlice{""},
			search: "",
			want:   true,
		},
	}
}

func TestStringSlice_IsIn(t *testing.T) {
	tests := getIsInTestCases()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.given.IsIn(test.search)
			if test.want != actual {
				t.Errorf("expected %t but got %t", test.want, actual)
			}
		})
	}
}

func TestStringSlice_IsNotIn(t *testing.T) {
	tests := getIsInTestCases()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.given.IsNotIn(test.search)
			if !test.want != actual {
				t.Errorf("expected %t but got %t", test.want, actual)
			}
		})
	}
}

func Test(t *testing.T) {
	tests := []struct {
		name  string
		given slice.StringSlice
		diff  slice.StringSlice
		want  slice.StringSlice
	}{
		{
			name:  "empty given, empty diff",
			given: slice.StringSlice{},
			diff:  slice.StringSlice{},
			want:  slice.StringSlice{},
		},

		{
			name:  "given with one element, empty diff",
			given: slice.StringSlice{"one"},
			diff:  slice.StringSlice{},
			want:  slice.StringSlice{"one"},
		},
		{
			name:  "given with two elements, empty diff",
			given: slice.StringSlice{"one", "two"},
			diff:  slice.StringSlice{},
			want:  slice.StringSlice{"one", "two"},
		},
		{
			name:  "empty given, diff with one element",
			given: slice.StringSlice{},
			diff:  slice.StringSlice{"one"},
			want:  slice.StringSlice{},
		},
		{
			name:  "empty given, diff with two elements",
			given: slice.StringSlice{},
			diff:  slice.StringSlice{"one", "two"},
			want:  slice.StringSlice{},
		},
		{
			name:  "given with one element, diff with one element - match",
			given: slice.StringSlice{"one"},
			diff:  slice.StringSlice{"one"},
			want:  slice.StringSlice{},
		},
		{
			name:  "given with one element, diff with one element - no match",
			given: slice.StringSlice{"one"},
			diff:  slice.StringSlice{"something"},
			want:  slice.StringSlice{"one"},
		},
		{
			name:  "given with two elements, diff with two elements - match complete",
			given: slice.StringSlice{"one", "two"},
			diff:  slice.StringSlice{"one", "two"},
			want:  slice.StringSlice{},
		},
		{
			name:  "given with two elements, diff with two elements - match only one",
			given: slice.StringSlice{"one", "two"},
			diff:  slice.StringSlice{"one", "three"},
			want:  slice.StringSlice{"two"},
		},
		{
			name:  "given with one element, diff with two elements - match complete",
			given: slice.StringSlice{"one"},
			diff:  slice.StringSlice{"two", "one"},
			want:  slice.StringSlice{},
		},
		{
			name:  "given with one element, diff with two elements - no match",
			given: slice.StringSlice{"one"},
			diff:  slice.StringSlice{"two", "three"},
			want:  slice.StringSlice{"one"},
		},
		{
			name:  "given with two elements, diff with one element - match",
			given: slice.StringSlice{"one", "two"},
			diff:  slice.StringSlice{"one"},
			want:  slice.StringSlice{"two"},
		},
		{
			name:  "given with two elements, diff with one element - no match",
			given: slice.StringSlice{"one", "two"},
			diff:  slice.StringSlice{"something"},
			want:  slice.StringSlice{"one", "two"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.given.Diff(test.diff)
			if len(test.want) != len(actual) {
				t.Errorf("exptect %v but got %v", test.want, actual)
			}

			for _, v := range test.want {
				if actual.IsNotIn(v) {
					t.Errorf("exptect %v but got %v", test.want, actual)
				}
			}
		})
	}
}
