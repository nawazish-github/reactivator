//Package observable represents the source of stream or sequence of data.
//It emits data, errors or end of observable signal to its subscribers.
package observable

import (
	"errors"
	"log"

	"github.com/nawazish-github/reactivator/observer"
)

type Observable <-chan interface{}

func From(data interface{}) (Observable, error) {
	switch t := data.(type) {
	case []interface{}:
		obsChan := make(chan interface{}, len(t))
		go func() {
			defer close(obsChan)
			for _, val := range t {
				obsChan <- val
			}
		}()
		return obsChan, nil
	case <-chan interface{}:
		return Observable(t), nil
	case chan interface{}:
		return Observable(t), nil
	default:
		return nil, errors.New("Invalid type. Could not create Observable from data!")
	}
}

func (observable Observable) Subscribe(observer observer.Observer) chan struct{} {
	subscriptionChannel := make(chan struct{})
	go func() {
		var e error = nil
		if observable != nil {
			for seq := range observable {
				switch tipe := seq.(type) {
				case error:
					observer.OnError(tipe)
					e = tipe
				default:
					observer.OnNext(tipe)
				}
			}
			if e == nil {
				observer.OnComplete()
			} else {
				log.Println("observable did not complete successfully")
			}
			subscriptionChannel <- struct{}{}
		} else {
			log.Println("observable is nil")
		}
	}()
	return subscriptionChannel
}
