package handlers

type EventHandler interface {
	Handle(seq interface{})
}
