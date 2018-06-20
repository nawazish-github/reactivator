package operators

import (
	"errors"

	"github.com/nawazish-github/reactivator/observable"
)

func Range(start int, length int) (observable.Observable, error) {

	if length < 0 {
		return nil, errors.New("Illegal Argument: length can't be negative with Range operator")
	}

	if length == 0 {
		return nil, nil
	}
	return nil, nil
}
