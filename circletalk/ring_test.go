package main

import "testing"

func TestElementRingUnconnected(t *testing.T) {
	const els = 10
	r := MakeRing(els)
	expected := make([]int, els)
	for i := 0; i < els; i++ {
		expected[i] = i
		if r.Elements()[i].Val() != i {
			t.Error("Invalid initial value in element", i, " - ", r.Elements()[i].Val())
		}
	}

	for i := 0; i < els; i++ {
		r.Step()
		expected = append(expected[els-1:], expected[:els-1]...)
		for i, e := range expected {
			if r.Elements()[i].Val() != e {
				t.Error("Value error for element", i, "expected", e, "got", r.Elements()[i].Val())
			}
		}

	}
	// after els iterations the values should be back in the correct order

	for i := 0; i < els; i++ {
		if r.Elements()[i].Val() != i {
			t.Error("Invalid initial value in element", i, " - ", r.Elements()[i].Val())
		}
	}

}
