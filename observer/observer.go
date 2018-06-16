package observer

type Observer struct {
	onNext     OnNext
	onError    OnError
	onComplete OnComplete
}

type OnNext func(interface{})
type OnError func(interface{})
type OnComplete func(interface{})

func (observer *Observer) Dispose() {

}
