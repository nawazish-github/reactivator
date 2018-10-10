package observable

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupByOperator_ShouldReturnValidNumberOfGroupsFromTheUnderlyingObservable(t *testing.T) {
	obs := initializeTestObservableOfMixedTypes()
	grpBy := obs.GroupBy()

	c := 0
	for items := range grpBy {
		_ = items
		c++
	}
	if c != 3 {
		t.Errorf("Did not return valid number of groups. Expected 3, got %d ", c)
	}
}

func TestGroupByOperator_ShouldGroupElementsCorrectly(t *testing.T) {
	obs := initializeTestObservableOfMixedTypes()
	grpBy := obs.GroupBy()
	var intArr []int
	var floatArr []float64
	var strArr []string

	for items := range grpBy {
		grp := items.([]interface{})

		for _, item := range grp {
			tipe := reflect.TypeOf(item)
			if tipe.Kind() == reflect.Int {
				intArr = append(intArr, item.(int))
				continue
			}
			if tipe.Kind() == reflect.Float64 {
				floatArr = append(floatArr, item.(float64))
				continue
			}
			if tipe.Kind() == reflect.String {
				strArr = append(strArr, item.(string))
				continue
			}
		}
	}

	assert.Equal(t, []int{11, 22, 33}, intArr, "Int group did not match")
	assert.Equal(t, []float64{3.2, 1.4}, floatArr, "float group did not match")
	assert.Equal(t, []string{"a"}, strArr, "string group did not match")
}
