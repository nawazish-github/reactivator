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
