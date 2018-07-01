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
	testObservable.Scan(nil, nil)
}

func TestScanOperator_ShouldReturnSameObservableIfScanFuncIsNil(t *testing.T) {

	var testObservable Observable
	obs, _ := testObservable.From([]interface{}{1, 2, 3})
	o := obs.Scan(nil, nil)
	if !CompareIntegerSlices(o, []interface{}{1, 2, 3}) {
		t.Errorf("Scan did not return the same Observable when input function was nil")
	}
}

func TestScanOperator_ShouldReturnValidObservable(t *testing.T) {
	var scanFunc ScanFunc = func(x, y interface{}) interface{} {
		xInt := x.(int)
		yInt := y.(int)
		return interface{}((xInt + yInt))
	}

	var testObservable Observable
	obs, _ := testObservable.From([]interface{}{4, 5, 6})

	o := obs.Scan(scanFunc, 0)
	if !CompareIntegerSlices(o, []interface{}{4, 9, 15}) {
		t.Errorf("Scan did not apply ScanFunc correctly")
	}
}
