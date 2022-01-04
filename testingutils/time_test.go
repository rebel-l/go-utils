package testingutils_test

import (
	"testing"
	"time"

	"github.com/rebel-l/go-utils/testingutils"
)

func TestTimeParse(t *testing.T) {
	t.Parallel()

	fixture, err := time.Parse("2006", "2022")
	if err != nil {
		t.Fatalf("failed to init fixture: %v", err)
	}

	testCases := []struct {
		name   string
		layout string
		value  string
		expect time.Time
	}{
		{
			name:   "wrong layout",
			layout: "2000",
			value:  "2022",
		},
		{
			name:   "not parsable value",
			layout: "2006",
			value:  "whatisthis?",
		},
		{
			name:   "success",
			layout: "2006",
			value:  "2022",
			expect: fixture,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := testingutils.TimeParse(tc.layout, tc.value)

			if !tc.expect.Equal(got) {
				t.Errorf("expected '%s' but got '%s'", tc.expect, got)
			}
		})
	}
}
