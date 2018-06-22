package observable

import (
	"testing"
	"time"
)

func TestIntervalOperator_ShouldReturnValidIntervalObserver(t *testing.T) {
	var testObservable Observable
	duration := time.Duration(time.Second)

	observable, err := testObservable.Interval(duration)
	if err != nil {
		t.Fail()
	}
	if observable == nil {
		t.Fail()
	}

	doneChan := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		doneChan <- struct{}{}
	}()

	counter := 0
	var sliceOfInt []int
ForLoop:
	for {
		select {
		case <-doneChan:
			break ForLoop
		case <-observable:
			sliceOfInt = append(sliceOfInt, counter)
			counter++
		}
	}
	if len(sliceOfInt) != 5 {
		t.Fail()
	}
}

func TestIntervalOperator_ShouldReturnErrorWhenDurationIsNegative(t *testing.T) {
	duration := -1
	var testObservable Observable
	_, err := testObservable.Interval(time.Duration(duration))
	if err == nil {
		t.Fail()
	}
}
