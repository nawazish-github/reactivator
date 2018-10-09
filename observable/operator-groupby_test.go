package observable

import (
	"testing"
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
