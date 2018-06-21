package observable

import (
	"errors"
	"testing"

	"github.com/nawazish-github/reactivator/handlers"
	"github.com/nawazish-github/reactivator/observer"
	"github.com/stretchr/testify/assert"
)

func TestFromOperatorWithUniformDataAndDefaultOnNextObserver(t *testing.T) {
	var testObservable Observable
	testTable := []interface{}{1, 2, 3}
	observable, err := testObservable.From(testTable)
	if err != nil {
		t.Errorf("Could not create observable from Uniform data slice ")
	}
	data := []interface{}{}
	observer := observer.RegisterHandlers(handlers.OnNext(func(seq interface{}) {
		switch tipe := seq.(type) {
		case int:
			data = append(data, tipe)
		}
	}))
	subChan := observable.Subscribe(observer)
	<-subChan
	expectation := []interface{}{1, 2, 3}

	assert.Len(t, data, 3, "Length of sequence emitted did not match")
	assert.Equal(t, expectation, data, "Sequence data emitted by Observable and received by Observer did not match")
}

func TestFromOperatorWithUniformDataAndDefaultOnCompleteObserver(t *testing.T) {
	var testObservable Observable
	testTable := []interface{}{1, 2, 3}
	observable, err := testObservable.From(testTable)
	if err != nil {
		t.Errorf("Could not create observable from Uniform data slice ")
	}
	data := []interface{}{}
	observer := observer.RegisterHandlers(handlers.OnComplete(func() {
		data = append(data, 1) //adding one to data slice signifying that OnComplete was called
	}))
	subChan := observable.Subscribe(observer)
	<-subChan

	assert.Len(t, data, 1, "Length of sequence emitted did not match")
}

func TestFromOperator_ShouldCallOnlyOnCompleteHandler(t *testing.T) {
	var testObservable Observable
	testTable := []interface{}{}
	observable, err := testObservable.From(testTable)
	if err != nil {
		t.Errorf("Could not create observable from Uniform data slice ")
	}
	ob := observer.RegisterHandlers(handlers.OnComplete(func() {}),
		handlers.OnNext(func(seq interface{}) { t.Fail() }),
		handlers.OnError(func(e error) { t.Fail() }),
	)

	<-observable.Subscribe(ob)
}

func TestFromOperator_ShouldCallOnlyOnErrorHandler(t *testing.T) {
	var testObservable Observable
	testTable := []interface{}{errors.New("Error")}
	observable, err := testObservable.From(testTable)
	if err != nil {
		t.Errorf("Could not create observable from Uniform data slice ")
	}

	ob := observer.RegisterHandlers(handlers.OnComplete(func() { t.Fail() }),
		handlers.OnNext(func(seq interface{}) { t.Fail() }),
		handlers.OnError(func(e error) {}),
	)

	<-observable.Subscribe(ob)
}

func TestFromOperator_ShouldNotCallOnErrorHandler(t *testing.T) {
	var testObservable Observable
	testTable := []interface{}{1}
	observable, err := testObservable.From(testTable)
	if err != nil {
		t.Errorf("Could not create observable from Uniform data slice ")
	}

	ob := observer.RegisterHandlers(handlers.OnComplete(func() {}),
		handlers.OnNext(func(seq interface{}) {}),
		handlers.OnError(func(e error) { t.Fail() }),
	)

	<-observable.Subscribe(ob)
}

func TestFromOperator_ShouldNotCallOnCompleteHandler(t *testing.T) {
	var testObservable Observable
	testTable := []interface{}{1, 2, errors.New("Error")}
	observable, err := testObservable.From(testTable)
	if err != nil {
		t.Errorf("Could not create observable from Uniform data slice ")
	}

	//this flags help assert that the corresponding handler functions were indeed called.
	onNextHandlerFlag := true
	onErrorHandlerFlag := true

	ob := observer.RegisterHandlers(handlers.OnComplete(func() { t.Fail() }),
		handlers.OnNext(func(seq interface{}) { onNextHandlerFlag = false }),
		handlers.OnError(func(e error) { onErrorHandlerFlag = false }),
	)

	<-observable.Subscribe(ob)

	if onNextHandlerFlag || onErrorHandlerFlag {
		t.Fail()
	}
}
