package assert

import (
	"reflect"
	"testing"
)

// Equal checks if values are equal
func Equal(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("Received %v (type %v), expected %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
	}
}

// DeepEqual checks if arrays/slices are equal
func DeepEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Received %v (type %v), expected %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
	}
}
