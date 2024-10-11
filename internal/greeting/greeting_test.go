package greeting

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Test")
	expected := "Hello, Test!"
	if result != expected {
		t.Errorf("Hello(\"Test\") = %q, want %q", result, expected)
	}
}