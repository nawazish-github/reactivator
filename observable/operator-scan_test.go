package observable

import (
	"testing"
)

func TestScanOperator_ShouldPanicIfObservableOnWhichScanIsCalledIsNil(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("Scan did not panic when Observable was nil")
		}
	}()
	var testObservable Observable
	testObservable.Scan(nil)
}

func TestScanOperator_ShouldReturnSameObservableIfScanFuncIsNil(t *testing.T) {

	var testObservable Observable
	obs, _ := testObservable.From([]interface{}{1, 2, 3})
	o := obs.Scan(nil)
	if !CompareIntegerSlices(o, []int{1, 2, 3}) {
		t.Errorf("Scan did not return the same Observable when input function was nil")
	}
}
