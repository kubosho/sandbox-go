package stack

import "testing"

func Test(t *testing.T) {
	s := &Stack{}

	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	s.Push("dataD")

	if s.Pop() != "dataD" {
		t.Error("Unexpected value")
	}

	if s.Pop() != "dataC" {
		t.Error("Unexpected value")
	}

	if s.Pop() != "dataB" {
		t.Error("Unexpected value")
	}

	if s.Pop() != "dataA" {
		t.Error("Unexpected value")
	}

	if s.Pop() != "" {
		t.Error("Unexpected value")
	}
}
