package signal

import (
	"testing"
)

// print messages
func TestHandler_print(t *testing.T) {
	t.Log("here is a plain message")
	t.Logf("here is my message with a number: %d", 123)
}

// fail test
func TestHandler_fail(t *testing.T) {
	t.Log("here is a plain message")
	t.Fail()
	t.Logf("here is my message with a number: %d", 123)
}

// fail test and stop execution
func TestHandler_failAndStop(t *testing.T) {
	t.Log("here is a plain message")
	t.FailNow()
	t.Logf("here is my message with a number: %d", 123)
}

// use error method
func TestHandler_error(t *testing.T) {
	t.Errorf("here is my message with a number: %d", 123)
}

// use fatal method
func TestHandler_fatal(t *testing.T) {
	t.Fatalf("here is my message with a string %s", "hello")
}
