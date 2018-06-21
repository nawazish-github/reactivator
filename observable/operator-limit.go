package observable

import "errors"

func (observable Observable) Limit(limit int) (Observable, error) {
	if observable == nil {
		return nil, errors.New("Nil Observable: Can't apply Limit operator on Nil Observer")
	}

	obsChan := make(chan interface{}, limit) //is buffering a right idea?

	counter := 1
	for val := range observable {
		if counter <= limit {
			counter++
			obsChan <- val
		} else {
			close(obsChan)
			break
		}
	}
	return Observable(obsChan), nil
}
