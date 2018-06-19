package observable

import (
	"testing"

	"github.com/nawazish-github/reactivator/handlers"
	"github.com/nawazish-github/reactivator/observer"
	"github.com/stretchr/testify/assert"
)

func TestFromOperatorWithUniformDataAndDefaultOnNextObserver(t *testing.T) {
	testTable := []interface{}{1, 2, 3}
	observable, err := From(testTable)
	if err != nil {
		t.Errorf("Could not create observable from Uniform data slice ")
	}
	data := []interface{}{}
	observer := observer.New(handlers.OnNext(func(seq interface{}) {
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
