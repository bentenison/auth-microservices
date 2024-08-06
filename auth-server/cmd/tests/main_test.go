package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup code if needed

	code := m.Run()
	// Teardown code if needed
	os.Exit(code)
}
