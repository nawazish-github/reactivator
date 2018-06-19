package observer

import (
	"log"

	"github.com/nawazish-github/reactivator/handlers"
)

type Observer struct {
	OnNext     handlers.OnNext
	OnError    handlers.OnError
	OnComplete handlers.OnComplete
}

func (observer *Observer) Dispose() {

}

func RegisterHandlers(eventHandlers ...handlers.EventHandler) Observer {
	observer := DefaultObserver
	if len(eventHandlers) > 0 {
		for _, handler := range eventHandlers {
			switch tipe := handler.(type) {
			case handlers.OnNext:
				observer.OnNext = tipe
			case handlers.OnComplete:
				observer.OnComplete = tipe
			case handlers.OnError:
				observer.OnError = tipe
			}
		}
	}
	return observer
}

func (observer Observer) Handle(seq interface{}) {
	switch tipe := seq.(type) {
	case error:
		observer.OnError(tipe)
	default:
		observer.OnNext(tipe)
	}
}

var DefaultObserver = Observer{
	OnNext: func(seq interface{}) {
		log.Println("Default OnNext: ", seq)
	},
	OnComplete: func() {
		log.Println("Default OnComplete signalled")
	},
	OnError: func(e error) {
		log.Fatal(e)
	},
}
