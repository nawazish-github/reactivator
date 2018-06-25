package observable

import (
	"testing"
)

func TestFlatMapOperator_ShouldPanicIfInputSliceLenIsZero(t *testing.T) {
	defer func() {
		if x := recover(); x == nil {
			t.Fail()
		}
	}()
	var testObservable Observable

	testObservable.FlatMap()
}
