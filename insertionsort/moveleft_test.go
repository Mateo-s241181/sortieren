package insertionsort

import (
	"sortieren/testhelpers"
	"testing"
)

func TestMoveLeft(t *testing.T) {
	list := []int{3, 0, 2, 1}
	MoveLeft(list, 3)
	testhelpers.AssertListsEqual(t, []int{3, 0, 1, 2}, list)
}
