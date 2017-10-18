package main

import "testing"

func TestFib(t *testing.T) {
	if fib(1) != 1 {
		t.Error("Unexpected fibonacci value")
	}

	if fib(10) != 55 {
		t.Error("Unexpected fibonacci value")
	}
}
