//Package observable represents the source of stream or sequence of data.
//It emits data, errors or end of observable signal to its subscribers.
package observable

type Observable <-chan interface{}
