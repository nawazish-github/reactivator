package operators

import (
	"errors"
	"time"

	"github.com/nawazish-github/reactivator/observable"
)

func Range(start int, length int) (observable.Observable, error) {

	if length < 0 {
		return nil, errors.New("Illegal Argument: length can't be negative with Range operator")
	}

	if length == 0 {
		return nil, nil
	}

	observable := make(chan interface{})
	go func() {
		defer close(observable)
		counter := 0
		for i := start; counter < length; counter++ {
			observable <- i
			i++
		}
	}()
	return observable, nil
}

func Interval(d time.Duration) (observable.Observable, error) {

	if d < 0 {
		return nil, errors.New("Illegal Argument: Negative duration for Interval operator")
	}

	observable := make(chan interface{})

	ticker := time.NewTicker(d)
	//defer ticker.Stop()

	counter := 0
	go func() {
		for {
			select {
			case <-ticker.C:
				observable <- counter
				counter++
			}
		}
	}()

	return observable, nil
}
