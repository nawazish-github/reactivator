package observable

import "errors"

func (observable Observable) Range(start int, length int) (Observable, error) {

	if length < 0 {
		return nil, errors.New("Illegal Argument: length can't be negative with Range operator")
	}

	if length == 0 {
		return nil, nil
	}

	obs := make(chan interface{})
	go func() {
		defer close(obs)
		counter := 0
		for i := start; counter < length; counter++ {
			obs <- i
			i++
		}
	}()
	return observable, nil
}
