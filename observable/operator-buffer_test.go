package observable

import (
	"testing"
)

func TestBufferOperator_ShouldPanicWhenBufferSizeIsZero(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Buffer did not panic when buffer size was zero")
		}
	}()

	ob := initializeTestObservableWithIntegers()
	ob.Buffer(0)
}

func TestBufferOperator_ShouldReturnValidObservable(t *testing.T) {
	ob := initializeTestObservableWithIntegers()
	buffObservable := ob.Buffer(3)
	var result []interface{}

	for list := range buffObservable {
		result = append(result, list)
	}

	if len(result) != 2 {
		t.Fail()
	}

	for _, r := range result {
		d := r.([]interface{})
		if len(d) != 3 {
			t.Fail()
		}
	}
}
