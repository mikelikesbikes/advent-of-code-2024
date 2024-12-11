package testutil

import (
	"reflect"
	"testing"
)

func AssertEqual(t *testing.T, expected any, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func Assert(t *testing.T, expected bool) {
	t.Helper()
	if !expected {
		t.Errorf(`expected true, got false`)
	}
}
