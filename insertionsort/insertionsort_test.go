package insertionsort

import (
	"sortieren/testhelpers"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	list := []int{3, 2, 1, 7, 9, 5}
	InsertionSort(list)
	testhelpers.AssertListsEqual(t, []int{1, 2, 3, 5, 7, 9}, list)
}
