package main

import "testing"

func TestWorking(t *testing.T) {
	result := Working()
	if result != "Working!" {
		t.Errorf("Invalid!")
	}
}
