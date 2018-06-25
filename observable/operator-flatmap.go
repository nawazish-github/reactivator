package observable

import (
	"sync"
)

func (observables Observable) FlatMap() Observable {

	if observables == nil {
		panic("Nil Pointer: FlatMap called on Nil Observable")
	}
	var wg sync.WaitGroup
	var sliceOfObservables []Observable
	obs := make(chan interface{})

	for ob := range observables {
		x := ob.(Observable)
		sliceOfObservables = append(sliceOfObservables, x)
	}

	wg.Add(len(sliceOfObservables))

	go func() {
		defer close(obs)
		wg.Wait()
	}()

	for _, observable := range sliceOfObservables {
		go func(observable Observable) {
			defer wg.Done()
			for val := range observable {
				obs <- val
			}
		}(observable)
	}

	return Observable(obs)
}
