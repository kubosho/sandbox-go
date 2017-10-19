package stack

import "testing"

func Test(t *testing.T) {
	s := &Stack{limit: 2}

	s.Push("dataA")
	s.Push("dataB")
	s.Push("dataC")
	s.Push("dataD")

	var actual1 = s.Pop()
	if actual1 != "dataD" {
		t.Errorf("Unexpected value %v", actual1)
	}

	var actual2 = s.Pop()
	if actual2 != "dataC" {
		t.Errorf("Unexpected value %v", actual2)
	}

	var actual3 = s.Pop()
	if actual3 != "" {
		t.Errorf("Unexpected value %v", actual3)
	}
}
