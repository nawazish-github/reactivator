package observable

type ScanFunc func(x, y interface{}) interface{}

func (observable Observable) Scan(fx ScanFunc, zeroValue interface{}) Observable {
	if observable == nil {
		panic("Illegal Argument: Scan called on nil Observable")
	}

	if fx == nil {
		return observable
	}

	scanChan := make(chan interface{})
	go func() {
		var acc interface{} = zeroValue
		defer close(scanChan)
		for val := range observable {
			acc = fx(val, acc)
			scanChan <- acc
		}
	}()
	return Observable(scanChan)
}
