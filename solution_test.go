package solution

import (
	"reflect"
	"testing"
)

func TestChange(t *testing.T) {
	r := Change(99, 100)
	expected := map[float64]int{
		1: 1,
	}

	if !reflect.DeepEqual(r, expected) {
		t.Errorf("expected %#v\nbut got %#v\n", expected, r)
	}
}
