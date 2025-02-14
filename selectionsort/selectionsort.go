package selectionsort

import "sortieren/testhelpers"

// SelectionSort sortiert eine Liste von Zahlen mit dem SelectionSort-Algorithmus.
func SelectionSort(list []int) {

	if len(list) < 2 {
		return
	}

	//Sucht nach dem kleinsten Element der Liste
	pos := SmallestPos(list)
	//Schreibt es nach vorne
	testhelpers.SwapElements(list, 0, pos)
	//Wiederholen ohne kleinstes Element
	SelectionSort(list[1:])
}
