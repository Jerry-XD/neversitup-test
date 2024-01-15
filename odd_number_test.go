package main

import (
	"testing"
)

func TestFindOddNumber(t *testing.T) {
	var resultCase1 = FindOddNumber([]int{7})
	var resultCase2 = FindOddNumber([]int{0})
	var resultCase3 = FindOddNumber([]int{1, 1, 2})
	var resultCase4 = FindOddNumber([]int{0, 1, 0, 1, 0})
	var resultCase5 = FindOddNumber([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})

	var expectResultCase1 = 7
	var expectResultCase2 = 0
	var expectResultCase3 = 2
	var expectResultCase4 = 0
	var expectResultCase5 = 4

	if resultCase1 != expectResultCase1 {
		t.Errorf("case1 error : have %d, want %d", resultCase1, expectResultCase1)
	}

	if resultCase2 != expectResultCase2 {
		t.Errorf("case2 error : have %d, want %d", resultCase2, expectResultCase2)
	}

	if resultCase3 != expectResultCase3 {
		t.Errorf("case3 error : have %d, want %d", resultCase3, expectResultCase3)
	}

	if resultCase4 != expectResultCase4 {
		t.Errorf("case4 error : have %d, want %d", resultCase4, expectResultCase4)
	}

	if resultCase5 != expectResultCase5 {
		t.Errorf("case5 error : have %d, want %d", resultCase5, expectResultCase5)
	}
}
