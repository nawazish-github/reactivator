package observable

import (
	"testing"
)

func TestLimitOperator_ShouldReturnErrorWhenLengthIsNegative(t *testing.T) {
	length := -1
	var testObservable Observable

	_, err := testObservable.Limit(length)

	if err == nil {
		t.Fail()
	}
}

func TestLimitOperator_ShouldReturnErrorWhenObservableIsNil(t *testing.T) {
	length := 6
	var testObservable Observable

	_, err := testObservable.Limit(length)

	if err == nil {
		t.Fail()
	}
}
