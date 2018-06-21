package observable

import (
	"errors"
	"time"
)

func (observable Observable) Interval(d time.Duration) (Observable, error) {

	if d < 0 {
		return nil, errors.New("Illegal Argument: Negative duration for Interval operator")
	}

	obs := make(chan interface{})

	ticker := time.NewTicker(d)
	//defer ticker.Stop()

	counter := 0
	go func() {
		for {
			select {
			case <-ticker.C:
				obs <- counter
				counter++
			}
		}
	}()

	return observable, nil
}
