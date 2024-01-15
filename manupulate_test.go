package main

import (
	"slices"
	"testing"
)

func TestManipulate(t *testing.T) {
	var resultCase1 = Manipulate("a")
	var resultCase2 = Manipulate("ab")
	var resultCase3 = Manipulate("abc")
	var resultCase4 = Manipulate("aabb")

	var expectResultCase1 = []string{"a"}
	var expectResultCase2 = []string{"ab", "ba"}
	var expectResultCase3 = []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	var expectResultCase4 = []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}

	if !slices.Equal(resultCase1, expectResultCase1) {
		t.Errorf("case1 error : have %s, want %s", resultCase1, expectResultCase1)
	}

	if !slices.Equal(resultCase2, expectResultCase2) {
		t.Errorf("case2 error : have %s, want %s", resultCase2, expectResultCase2)
	}

	if !slices.Equal(resultCase3, expectResultCase3) {
		t.Errorf("case3 error : have %s, want %s", resultCase3, expectResultCase3)
	}

	if !slices.Equal(resultCase4, expectResultCase4) {
		t.Errorf("case4 error : have %s, want %s", resultCase4, expectResultCase4)
	}
}
