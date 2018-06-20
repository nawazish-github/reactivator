package operators

import (
	"testing"
)

func TestRange_ShouldReturnErrorIfLengthIsNegative(t *testing.T) {
	start := 0
	length := -1
	_, err := Range(start, length)

	if err == nil {
		t.Fail()
	}
}
