package observable

import "reflect"

//GroupBy groups the sequence of observables by their concrete types; this is the default behavior
//GroupBy returns an Observable in which each element is a slice of grouped elements.
func (observable Observable) GroupBy() Observable {
	obsChan := make(chan interface{})

	go func() {
		defer close(obsChan)
		groups := make(map[reflect.Type][]interface{})
		flag := false
		for item := range observable {
			flag = false
			t := reflect.TypeOf(item)
			for key, value := range groups {
				if key == t {
					value = append(value, item)
					groups[key] = value
					flag = true
					break
				}
			}
			if flag != true {
				var sl []interface{}
				sl = append(sl, item)
				groups[t] = sl
			}
		}
		for _, value := range groups {
			obsChan <- value
		}
	}()

	return Observable(obsChan)
}
