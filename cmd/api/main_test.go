package main

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "value")
	defer os.Unsetenv("TEST_KEY")

	if getEnv("TEST_KEY", "fallback") != "value" {
		t.Error("Expected value")
	}
	if getEnv("NON_EXISTENT", "fallback") != "fallback" {
		t.Error("Expected fallback")
	}
}
