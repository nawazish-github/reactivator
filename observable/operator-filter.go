package observable

import (
	"errors"
)

type filterFunc func(x interface{}) bool

func (observable Observable) Filter(fx filterFunc) (Observable, error) {

	if observable == nil {
		return nil, errors.New("Nil Observable: Filter operator called on Nil Observable")
	}

	if fx == nil {
		return observable, nil
	}
	return nil, nil
}
