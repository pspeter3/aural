package aural_test

import (
	"testing"
)

func assert(t *testing.T, condition bool, message string) {
	if !condition {
		t.Fatal(message)
	}
}

func ok(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func equals(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Fatal("%s does not equal %s", actual, expected)
	}
}
