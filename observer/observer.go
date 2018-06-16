package observer

type Observer struct {
	OnNext     OnNext
	OnError    OnError
	OnComplete OnComplete
}

type OnNext func(interface{})
type OnError func(error)
type OnComplete func()

func (observer *Observer) Dispose() {

}
