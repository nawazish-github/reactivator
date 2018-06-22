package observable

import (
	"testing"
)

func TestFilterOperator_ShouldReturnErrorWhenObservableOnWhichFilterIsCalledIsNil(t *testing.T) {
	var testObservable Observable
	_, err := testObservable.Filter(func(x interface{}) bool {
		return false
	})

	if err == nil {
		t.Fail()
	}
}
