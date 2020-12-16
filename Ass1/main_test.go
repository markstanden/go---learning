package main

import "testing"

func TestCreateSlice(t *testing.T) {
	size := 100
	var source = newIntSlice(size)
	if len(source) != size+1 {
		t.Errorf("Expected slice of length %v, but got %v", size+1, len(source))
	}
}
