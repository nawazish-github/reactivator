package observable

import (
	"testing"
)

func TestMapOperator_ShouldReturnTheSameObservableIfMapperFunctionIsNil(t *testing.T) {
	obsChan := make(chan interface{}, 5)
	for i := 1; i <= 5; i++ {
		obsChan <- i
	}
	close(obsChan)
	testObservable := Observable(obsChan)

	obs, err := testObservable.Map(nil)

	if err != nil {
		t.Fail()
	}

	if obs == nil {
		t.Fail()
	}

	expectation := []int{1, 2, 3, 4, 5}
	var result []int

	for val := range obs {
		i := val.(int)
		result = append(result, i)
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

func TestMapOperator_ShouldReturnErrorIfObservableOnWhichMapOperatorIsCalledIsNil(t *testing.T) {
	var observable Observable
	mapper := func(x interface{}) interface{} {
		return x
	}
	_, err := observable.Map(mapper)

	if err == nil {
		t.Fail()
	}
}

func TestMapOperator_ShouldReturnValidObservableAfterApplyingMapperFunction(t *testing.T) {
	obsChan := make(chan interface{}, 5)
	for i := 1; i <= 5; i++ {
		obsChan <- i
	}
	close(obsChan)
	testObservable := Observable(obsChan)
	expectation := []int{3, 4, 5, 6, 7}
	var result []int
	obs, err := testObservable.Map(func(x interface{}) interface{} {
		y := x.(int)
		return y + 2
	})

	if err != nil {
		t.Fail()
	}

	if obs == nil {
		t.Fail()
	}

	for val := range obs {
		x := val.(int)
		result = append(result, x)
	}

	if len(expectation) != len(result) {
		t.Fail()
	}

	for index := range result {
		if expectation[index] != result[index] {
			t.Fail()
		}
	}
}
