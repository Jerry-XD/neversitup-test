package main

import "testing"

func TestCountSmilyFace(t *testing.T) {
	var resultCase1 = CountSmilyFace([]string{":)", ";(", ";}", ":-D"})
	var resultCase2 = CountSmilyFace([]string{";D", ":-(", ":-)", ";~)"})
	var resultCase3 = CountSmilyFace([]string{";]", ":[", ";*", ":$", ";-D"})

	var expectResultCase1 = 2
	var expectResultCase2 = 3
	var expectResultCase3 = 1

	if resultCase1 != expectResultCase1 {
		t.Errorf("case1 error : have %d, want %d", resultCase1, expectResultCase1)
	}

	if resultCase2 != expectResultCase2 {
		t.Errorf("case2 error : have %d, want %d", resultCase2, expectResultCase2)
	}

	if resultCase3 != expectResultCase3 {
		t.Errorf("case3 error : have %d, want %d", resultCase3, expectResultCase3)
	}

}
