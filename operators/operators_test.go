package operators

import (
	"testing"
	"time"
)

func TestRange_ShouldReturnErrorIfLengthIsNegative(t *testing.T) {
	start := 0
	length := -1
	_, err := Range(start, length)

	if err == nil {
		t.Fail()
	}
}

func TestRange_ShouldReturnNilObserverIfLengthIsZero(t *testing.T) {
	start := 0
	length := 0
	observable, _ := Range(start, length)

	if observable != nil {
		t.Fail()
	}
}

func TestRange_ShouldReturnValidObservable(t *testing.T) {
	start := 0
	length := 4

	observable, err := Range(start, length)

	if err != nil {
		t.Fail()
	}

	if observable == nil {
		t.Fail()
	}

	var si []int
	for value := range observable {
		if v, ok := value.(int); ok {
			si = append(si, v)
		}
	}
	if len(si) != 4 {
		t.Fail()
	}
	counter := start

	for i := range si {
		if counter != i {
			t.Fail()
		}
		counter++
	}
}

func TestIntervalOperator_ShouldReturnValidIntervalObserver(t *testing.T) {
	duration := time.Duration(time.Second)

	observable, err := Interval(duration)
	if err != nil {
		t.Fail()
	}
	if observable == nil {
		t.Fail()
	}

	doneChan := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Second)
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
	if len(sliceOfInt) != 10 {
		t.Fail()
	}
}
