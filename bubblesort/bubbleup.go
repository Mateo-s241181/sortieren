package bubblesort

import "sortieren/testhelpers"

// BubbleUp implementiert eine Iteration des Bubble-Sort-Algorithmus.
// Es iteriert über die Liste und vergleicht jedes Element mit dem nächsten.
// Wenn das Element größer als das nächste ist, werden sie vertauscht.
// Die Funktion gibt true zurück, wenn mindestens ein Tausch durchgeführt wurde.
func BubbleUp(list []int) bool {
	swapped := false

	for i := range list {

		//Wenn i am letzten Element algelangt ist soll er nicht mit dem nächsten vergleichen (OUT OF RANGE)
		if i == len(list)-1 {
			break
		}

		//Element mit dem nächsten verlgeichen
		if list[i] > list[i+1] {

			//Elemente vertauschen
			testhelpers.SwapElements(list, i, i+1)

			//Vertauaschung dokumentieren
			swapped = true
		}

	}

	//Wurde ein Element getauscht?
	return swapped
}
