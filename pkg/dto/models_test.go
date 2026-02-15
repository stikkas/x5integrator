package dto

import "testing"

func Test_types(t *testing.T) {
	ot := Status
	if int(ot) != 1 {
		t.Error("Status should be 1 instead of ", int(ot))
	}
	ot = Study
	if int(ot) != 2 {
		t.Error("Study should be 2 instead of ", int(ot))
	}
	ot = Subscribe
	if int(ot) != 3 {
		t.Error("Subscribe should be 3 instead of ", int(ot))
	}
	ot = Unsubscribe
	if int(ot) != 4 {
		t.Error("Unsubscribe should be 4 instead of ", int(ot))
	}

	st := Topic
	if int(st) != 1 {
		t.Error("Topic type should be 1 instead of ", int(st))
	}
	st = Track
	if int(st) != 2 {
		t.Error("Track type should be 2 instead of ", int(st))
	}
}
