//Package observable represents the source of stream or sequence of data.
//It emits data, errors or end of observable signal to its subscribers.
package observable

import (
	"log"

	"github.com/nawazish-github/reactivator/observer"
)

type Observable <-chan interface{}

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
