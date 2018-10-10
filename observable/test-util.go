package observable

func CompareIntegerSlices(o Observable, y []int) bool {
	var x []int

	for val := range o {
		v := val.(int)
		x = append(x, v)
	}

	if len(x) != len(y) {
		return false
	}

	if x == nil || y == nil {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func initializeTestObservableWithIntegers() Observable {
	var testObservable Observable
	ob, _ := testObservable.From([]interface{}{1, 2, 3, 4, 5, 6})
	return ob
}

func initializeTestObservableOfMixedTypes() Observable {
	var testObservable Observable
	ob, _ := testObservable.From([]interface{}{11, 22, 33, 3.2, 1.4, "a"})
	return ob
}
