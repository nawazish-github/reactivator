package observable

type ScanFunc func(x interface{}) interface{}

func (observable Observable) Scan(fx ScanFunc) Observable {
	if observable == nil {
		panic("Illegal Argument: Scan called on nil Observable")
	}
	return nil
}
