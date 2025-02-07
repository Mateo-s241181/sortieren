package bubblesort

// BubbleSort sortiert the gegebene Liste mit dem Bubble-Sort-Algorithmus.
func BubbleSort(list []int) {

	if len(list) > 1 {

		//BubbleUp() callen, solange es true returned
		for BubbleUp(list) {
		}
	}
}
