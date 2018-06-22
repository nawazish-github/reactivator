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

func TestFilterOperator_ShouldApplyFilterFunctionAndReturnValidObservable(t *testing.T) {
	expectation := []int{1, 3, 5}

	obsChan := make(chan interface{}, 5)
	for i := 1; i <= 5; i++ {
		obsChan <- i
	}
	close(obsChan)
	testObservable := Observable(obsChan)

	obs, err := testObservable.Filter(func(x interface{}) bool {
		y := x.(int)

		if y%2 != 0 {
			return true
		} else {
			return false
		}

	})
	if err != nil {
		t.Fail()
	}

	if obs == nil {
		t.Fail()
	}

	var result []int

	for val := range obs {
		y := val.(int)
		result = append(result, y)
	}

	if len(expectation) != len(result) {
		t.Fail()
	}

	for index := range result {
		if result[index] != expectation[index] {
			t.Fail()
		}
	}
}
