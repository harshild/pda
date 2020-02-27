package main

import "testing"

func TestOpen(t *testing.T) {
	got := open([]byte("{}"))
	if got != false {
		t.Errorf("Abs(-1) = %t; want false", got)
	}
}
