package handlers

type (
	OnNext     func(seq interface{})
	OnComplete func()
	OnError    func(e error)
)

func (onNext OnNext) Handle(seq interface{}) {
	switch tipe := seq.(type) {
	case error:
		return
	default:
		onNext(tipe)
	}
}
func (onError OnError) Handle(seq interface{}) {
	switch tipe := seq.(type) {
	case error:
		onError(tipe)
	default:
		return
	}

}
func (onComplete OnComplete) Handle(seq interface{}) {
	onComplete()
}
