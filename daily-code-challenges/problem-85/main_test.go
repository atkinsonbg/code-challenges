package main

import "testing"

//This problem was asked by Facebook.
//
//Given three 32-bit integers x, y, and b, return x if b is 1 and y if b is 0, using
//only mathematical or bit operations. You can assume b can only be 1 or 0.

func Test1(t *testing.T) {

	got := FindB(4, 6, 1)
	want := 4

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Test2(t *testing.T) {

	got := FindB(4, 6, 0)
	want := 6

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
