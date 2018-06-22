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
