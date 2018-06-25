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

func TestFlatMapOperator_ShouldReturnValidFlatMapObservable(t *testing.T) {
	var observable Observable
	var slice []int
	expectation := []int{1, 2, 3, 4, 5, 6}
	obs1, _ := observable.From([]interface{}{1, 2, 3})
	obs2, _ := observable.From([]interface{}{4, 5, 6})
	soOb := []interface{}{obs1, obs2}
	fromOb, _ := observable.From(soOb)
	flatMapObser := fromOb.FlatMap()
	for val := range flatMapObser {
		x := val.(int)
		slice = append(slice, x)
	}

	if len(slice) != len(expectation) {
		t.Fail()
	}

	for _, val := range slice {
		var found bool
		for _, expectedData := range expectation {
			if expectedData == val {
				found = true
				break
			}
		}
		if !found {
			t.Fail()
			break
		}
	}
}
