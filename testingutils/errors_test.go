package testingutils_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rebel-l/go-utils/testingutils"
	"github.com/rebel-l/go-utils/testingutils/mocks/testing_mock"
)

func TestErrorsCheck(t *testing.T) {
	t.Parallel()

	errSame := errors.New("same") // nolint:goerr113

	testCases := []struct {
		name          string
		actualError   error
		expectedError error
		errMsg        string
	}{
		{
			name: "both errors nil",
		},
		{
			name:          "both errors same",
			actualError:   errSame,
			expectedError: errSame,
		},
		{
			name:          "actual nil, expected set",
			expectedError: errors.New("expected"), // nolint:goerr113
			errMsg:        "expected error 'expected' but got '<nil>'",
		},
		{
			name:        "actual set, expected nil",
			actualError: errors.New("actual"), // nolint:goerr113
			errMsg:      "expected error '<nil>' but got 'actual'",
		},
		{
			name:          "actual different than expected",
			actualError:   errors.New("actual"),   // nolint:goerr113
			expectedError: errors.New("expected"), // nolint:goerr113
			errMsg:        "expected error 'expected' but got 'actual'",
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockTesting := testing_mock.NewMockTB(ctrl)
			mockTesting.EXPECT().Helper().Times(1)

			if tc.errMsg != "" {
				mockTesting.EXPECT().Error(gomock.Eq(tc.errMsg)).Times(1)
			} else {
				mockTesting.EXPECT().Error(gomock.Any()).Times(0)
			}

			testingutils.ErrorsCheck(mockTesting, tc.expectedError, tc.actualError)
		})
	}
}
