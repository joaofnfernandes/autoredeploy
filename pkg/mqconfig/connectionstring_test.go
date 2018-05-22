package mqconfig

import (
	"testing"
)

var (
	validConnStr   = ConnectionString{"user", "pass", "localhost", 1234, "amqp"}
	invalidConnStr = ConnectionString{"", "pass", "localhost", 1234, "http"}
)

func TestValidate(t *testing.T) {
	err := validConnStr.validate()
	if err != nil {
		t.Fatalf("Validate says connection string is invalid, when it is valid")
	}
	err = invalidConnStr.validate()
	if err == nil {
		t.Fatalf("Validate says connection string is valid, when is its not")
	}
}

func TestString(t *testing.T) {
	expected := "amqp://user:pass@localhost:1234"
	got := validConnStr.String()
	if expected != got {
		t.Fatalf("Invalid connection string. Expected: %s, got: %s", expected, got)
	}
}
