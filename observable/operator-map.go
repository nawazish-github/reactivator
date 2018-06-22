package observable

import "errors"

type mapFunction func(x interface{}) interface{}

func (observable Observable) Map(fx mapFunction) (Observable, error) {

	if observable == nil {
		return nil, errors.New("Nil observable: Map operator called on nil observable")
	}

	if fx == nil {
		return observable, nil
	}

	mapChan := make(chan interface{})
	go func() {
		for val := range observable {
			y := fx(val)
			mapChan <- y
		}
		close(mapChan)
	}()

	return Observable(mapChan), nil
}
