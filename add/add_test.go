package add

import "testing"

func TestAdd1(t *testing.T) {
	if Add(1, 2) == 3 {
		t.Log("passed")
	}
}

func TestAdd2(t *testing.T) {
	if Add(1, 1) == 3 {
		t.Error("error")
	}
}
