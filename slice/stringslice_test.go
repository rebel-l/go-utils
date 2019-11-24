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

func TestStringSlice_Len(t *testing.T) {
	testCases := []struct {
		name  string
		given slice.StringSlice
		want  int
	}{
		{
			name: "no elements",
			want: 0,
		},
		{
			name:  "one elements",
			given: slice.StringSlice{"one"},
			want:  1,
		},
		{
			name:  "two elements",
			given: slice.StringSlice{"one", "two"},
			want:  2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.given.Len() != testCase.want {
				t.Errorf("expected length is %d but got %d", testCase.given.Len(), testCase.want)
			}
		})
	}
}

func TestStringSlice_IsEqual(t *testing.T) {
	testCases := []struct {
		name string
		a, b slice.StringSlice
		want bool
	}{
		{
			name: "first empty",
			b:    slice.StringSlice{"one"},
			want: false,
		},
		{
			name: "second empty",
			a:    slice.StringSlice{"one"},
			want: false,
		},
		{
			name: "first has more items",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one"},
			want: false,
		},
		{
			name: "second has more items",
			a:    slice.StringSlice{"one"},
			b:    slice.StringSlice{"one", "two"},
			want: false,
		},
		{
			name: "both has one item items",
			a:    slice.StringSlice{"one"},
			b:    slice.StringSlice{"one"},
			want: true,
		},
		{
			name: "both has two items in exactly same order",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one", "two"},
			want: true,
		},
		{
			name: "both has two items in different order",
			a:    slice.StringSlice{"two", "one"},
			b:    slice.StringSlice{"one", "two"},
			want: true,
		},
		{
			name: "slices has both two items but differs in values",
			a:    slice.StringSlice{"one", "three"},
			b:    slice.StringSlice{"one", "two"},
			want: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			given := testCase.a.IsEqual(testCase.b)
			if testCase.want != given {
				t.Errorf("slices are not equal, a:'%v' | b: '%v'", testCase.a, testCase.b)
			}
		})
	}
}

func BenchmarkStringSlice_IsEqual(b *testing.B) {
	cases := []struct {
		name string
		a, b slice.StringSlice
	}{
		{
			name: "differs in length",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one"},
		},
		{
			name: "differs in values",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one", "three"},
		},
		{
			name: "values are same and have same order",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one", "two"},
		},
		{
			name: "values are same but different order",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"two", "one"},
		},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = c.a.IsEqual(c.b)
			}
		})
	}
}

func TestStringSlice_IsSame(t *testing.T) {
	testCases := []struct {
		name string
		a, b slice.StringSlice
		want bool
	}{
		{
			name: "first empty",
			b:    slice.StringSlice{"one"},
			want: false,
		},
		{
			name: "second empty",
			a:    slice.StringSlice{"one"},
			want: false,
		},
		{
			name: "first has more items",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one"},
			want: false,
		},
		{
			name: "second has more items",
			a:    slice.StringSlice{"one"},
			b:    slice.StringSlice{"one", "two"},
			want: false,
		},
		{
			name: "both has one item items",
			a:    slice.StringSlice{"one"},
			b:    slice.StringSlice{"one"},
			want: true,
		},
		{
			name: "both has two items in exactly same order",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one", "two"},
			want: true,
		},
		{
			name: "both has two items in different order",
			a:    slice.StringSlice{"two", "one"},
			b:    slice.StringSlice{"one", "two"},
			want: false,
		},
		{
			name: "slices has both two items but differs in values",
			a:    slice.StringSlice{"one", "three"},
			b:    slice.StringSlice{"one", "two"},
			want: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			given := testCase.a.IsSame(testCase.b)
			if testCase.want != given {
				t.Errorf("slices are not same, a:'%v' | b: '%v'", testCase.a, testCase.b)
			}
		})
	}
}

func BenchmarkStringSlice_IsSame(b *testing.B) {
	cases := []struct {
		name string
		a, b slice.StringSlice
	}{
		{
			name: "differs in length",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one"},
		},
		{
			name: "differs in values",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one", "three"},
		},
		{
			name: "values are same and have same order",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"one", "two"},
		},
		{
			name: "values are same but different order",
			a:    slice.StringSlice{"one", "two"},
			b:    slice.StringSlice{"two", "one"},
		},
	}

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = c.a.IsSame(c.b)
			}
		})
	}
}
