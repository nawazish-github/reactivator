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

func TestLimitOperator_ShouldReturnValidObservableAfterApplyingLimit(t *testing.T) {

	observable := make(chan interface{}, 5)
	for i := 1; i <= 5; i++ {
		observable <- i
	}
	testObservable := Observable(observable)
	expectation := []interface{}{1, 2}

	length := 2
	obs, _ := testObservable.Limit(length)
	var result []interface{}

	if obs == nil {
		t.Fail()
	}

	for val := range obs {
		result = append(result, val)
	}

	if len(expectation) != len(result) {
		t.Fail()
	}

	for index := range expectation {
		if expectation[index] != result[index] {
			t.Fail()
		}
	}
}
