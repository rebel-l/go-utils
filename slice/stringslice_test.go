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
	t.Parallel()

	tests := getIsInTestCases()

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.given.IsIn(tc.search)
			if tc.want != actual {
				t.Errorf("expected %t but got %t", tc.want, actual)
			}
		})
	}
}

func TestStringSlice_IsNotIn(t *testing.T) {
	t.Parallel()

	tests := getIsInTestCases()

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.given.IsNotIn(tc.search)
			if !tc.want != actual {
				t.Errorf("expected %t but got %t", tc.want, actual)
			}
		})
	}
}

func Test(t *testing.T) {
	t.Parallel()

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
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.given.Diff(tc.diff)
			if len(tc.want) != len(actual) {
				t.Errorf("exptect %v but got %v", tc.want, actual)
			}

			for _, v := range tc.want {
				if actual.IsNotIn(v) {
					t.Errorf("exptect %v but got %v", tc.want, actual)
				}
			}
		})
	}
}

func TestStringSlice_Len(t *testing.T) {
	t.Parallel()

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
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if tc.given.Len() != tc.want {
				t.Errorf("expected length is %d but got %d", tc.given.Len(), tc.want)
			}
		})
	}
}

func TestStringSlice_IsEqual(t *testing.T) {
	t.Parallel()

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
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			given := tc.a.IsEqual(tc.b)
			if tc.want != given {
				t.Errorf("slices are not equal, a:'%v' | b: '%v'", tc.a, tc.b)
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
	t.Parallel()

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
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			given := tc.a.IsSame(tc.b)
			if tc.want != given {
				t.Errorf("slices are not same, a:'%v' | b: '%v'", tc.a, tc.b)
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

func TestStringSlice_String(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		data     slice.StringSlice
		expected string
	}{
		{
			name:     "empty slice",
			expected: "",
		},
		{
			name:     "one element",
			data:     []string{"one"},
			expected: "one",
		},
		{
			name:     "two elements",
			data:     []string{"one", "two"},
			expected: "one,two",
		},
		{
			name:     "three elements (second is empty string)",
			data:     []string{"one", "", "three"},
			expected: "one,,three",
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.data.String()
			if got != tc.expected {
				t.Errorf("expected '%s' but got '%s'", tc.expected, got)
			}
		})
	}
}

func TestStringSlice_Join(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name      string
		data      slice.StringSlice
		separator string
		expected  string
	}{
		{
			name:      "empty slice",
			separator: "<>",
			expected:  "",
		},
		{
			name:      "one element",
			data:      []string{"one"},
			separator: "-",
			expected:  "one",
		},
		{
			name:      "two elements",
			data:      []string{"one", "two"},
			separator: "|",
			expected:  "one|two",
		},
		{
			name:      "three elements (second is empty string)",
			data:      []string{"one", "", "three"},
			separator: ".",
			expected:  "one..three",
		},
	}

	for _, testCase := range testCases {
		tc := testCase

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.data.Join(tc.separator)
			if got != tc.expected {
				t.Errorf("expected '%s' but got '%s'", tc.expected, got)
			}
		})
	}
}
