package insertionsort

import "sortieren/testhelpers"

// MoveLeft bewegt das Element an der gegebenen Position nach links,
// bis es rechts von der nächstkleineren Zahl ist.
func MoveLeft(list []int, pos int) {

	swapped := false

	//Überprüfen ob Element an pos kleiner als Element an pos-1 ist
	// Argument pos > 0 MUSS an erster Stelle stehen, damit zuerst darauf überprüft wird (ansonsten OUT OF RANGE)
	if pos > 0 && list[pos] < list[pos-1] {

		//list[pos] wird eins nach links verschoben
		testhelpers.SwapElements(list, pos, pos-1)
		swapped = true
	}

	if swapped {
		MoveLeft(list, pos-1)
	}
}
