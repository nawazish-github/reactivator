package observable

import (
	"errors"
)

func (observable Observable) Limit(limit int) (Observable, error) {

	if limit < 0 {
		return nil, errors.New("Illegal Argument: Negative limit length")
	}

	if observable == nil {
		return nil, errors.New("Nil Observable: Can't apply Limit operator on Nil Observer")
	}

	obsChan := make(chan interface{}, limit) //is buffering a right idea?
	isObsChanClosed := false

	counter := 1
	for val := range observable {
		if counter <= limit {
			counter++
			obsChan <- val
		} else {
			close(obsChan)
			isObsChanClosed = true
			break
		}
	}
	if !isObsChanClosed {
		close(obsChan)
	}
	return Observable(obsChan), nil
}
