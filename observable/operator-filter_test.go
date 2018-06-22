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

func TestFilterOperator_ShouldReturnSameObservableWhenFilterFunctionIsNil(t *testing.T) {

	expectation := []int{1, 2, 3, 4, 5}

	obsChan := make(chan interface{}, 5)
	for i := 1; i <= 5; i++ {
		obsChan <- i
	}
	close(obsChan)
	testObservable := Observable(obsChan)

	obs, err := testObservable.Filter(nil)
	if err != nil {
		t.Fail()
	}
	var result []int
	for val := range obs {
		x := val.(int)
		result = append(result, x)
	}

	if len(result) != len(expectation) {
		t.Fail()
	}

	for index := range result {
		if result[index] != expectation[index] {
			t.Fail()
		}
	}
}
