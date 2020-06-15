//go:generate mockgen -destination=./mocks/testing_mock/tb_mock.go -package=testing_mock testing TB

package testingutils_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rebel-l/go-utils/testingutils/mocks/testing_mock"

	"github.com/google/uuid"

	"github.com/rebel-l/go-utils/testingutils"
)

func TestParse(t *testing.T) {
	want, err := uuid.NewRandom()
	if err != nil {
		t.Fatalf("failed to create UUID: %v", err)
	}

	got := testingutils.UUIDParse(t, want.String())

	if want.String() != got.String() {
		t.Errorf("wanted UUID '%s' but got '%s'", want.String(), got.String())
	}
}

func TestParse_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	u := "invalid UUID"

	mockTesting := testing_mock.NewMockTB(ctrl)
	mockTesting.EXPECT().Helper().Times(1)
	mockTesting.EXPECT().Fatal(
		gomock.Eq(
			fmt.Sprintf("failed to parse UUID from '%s': %v", u, "invalid UUID length: 12"),
		)).Times(1)

	_ = testingutils.UUIDParse(mockTesting, u)
}
