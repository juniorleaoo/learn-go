package average

import "testing"

func TestAverage(t *testing.T) {
	expected := 7.28
	actual := Average(7.2, 9.9, 6.1, 5.9)

	if actual != expected {
		t.Errorf("Expected %v, Actual %v", expected, actual)
	}
}
