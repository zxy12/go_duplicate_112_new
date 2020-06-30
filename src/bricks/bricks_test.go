package bricks

import (
	"testing"
)

func TestDoTest(t *testing.T) {
	if DoTest() != 1 {
		t.Error("test fail")
	}
}
