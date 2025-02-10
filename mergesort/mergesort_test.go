package mergesort

import (
	"sortieren/testhelpers"
	"testing"
)

func TestMergeSort(t *testing.T) {
	list := []int{112312, 2, 13, 12, 5, 6, 23, 25, 38, 77, 200, 1, 0}
	MergeSort(list)
	testhelpers.AssertListsEqual(t, []int{0, 1, 2, 5, 6, 12, 13, 23, 25, 38, 77, 200, 112312}, list)
}
