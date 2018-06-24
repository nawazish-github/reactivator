package observable

import (
	"testing"
)

func TestFlatMapOperator_ShouldPanicIfInputSliceLenIsZero(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Illegal Argument: soOb")
		}
	}()
	var testObservable Observable
	testObservable.FlatMap([]Observable{})
}
