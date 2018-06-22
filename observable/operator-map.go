package observable

type mapFunction func(x ...interface{}) interface{}

func (observable Observable) Map(fx mapFunction) (Observable, error) {

	if fx == nil {
		return observable, nil
	}
	return nil, nil
}
