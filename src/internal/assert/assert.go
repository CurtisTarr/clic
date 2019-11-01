package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("Received %v (type %v), expected %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
	}
}

func DeepEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Received %v (type %v), expected %v (type %v)", actual, reflect.TypeOf(actual), expected, reflect.TypeOf(expected))
	}
}
