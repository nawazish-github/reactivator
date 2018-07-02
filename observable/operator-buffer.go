package observable

func (o Observable) Buffer(buffSize int) Observable {

	if buffSize < 1 {
		panic("Buffer size cannot be less than one")
	}

	obs := make(chan interface{})

	go func() {
		defer close(obs)
		i := 1
		var slice []interface{}
		for val := range o {
			if i <= buffSize {
				slice = append(slice, val)
				i++
				if i > buffSize {
					obs <- slice
					i = 1
					slice = make([]interface{}, 0, 0)
				}
			}
		}
	}()

	return Observable(obs)
}
