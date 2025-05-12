package main_test

import (
	"bytes"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("Hello, World!")

	expected := "Hello, World!"
	if buf.String() != expected {
		t.Errorf("Expected %q but got %q", expected, buf.String())
	}
}
