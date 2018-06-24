package observable

func (observable Observable) FlatMap(soOb []Observable) Observable {

	if len(soOb) == 0 {
		panic("Illegal Argument: len of input slice to FlatMap is zero")
	}
	var obs Observable
	return obs
}
