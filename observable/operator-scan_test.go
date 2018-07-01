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
