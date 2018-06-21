package observable

import "errors"

func (observable Observable) From(data interface{}) (Observable, error) {
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
