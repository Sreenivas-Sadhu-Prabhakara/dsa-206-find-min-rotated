package main

import "testing"

func TestFindMin(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11}, // not rotated at all
		{[]int{2, 1}, 1},
		{[]int{1}, 1},
		{[]int{5, 1, 2, 3, 4}, 1}, // rotated by one
		{[]int{2, 3, 4, 5, 1}, 1}, // minimum at the very end
	}
	for _, c := range cases {
		if got := FindMin(c.in); got != c.want {
			t.Errorf("FindMin(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
