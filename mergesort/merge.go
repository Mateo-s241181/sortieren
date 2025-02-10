package mergesort

// Merge erwartet zwei sortierte Listen und gibt eine sortierte Liste zurück,
// die alle Elemente der beiden Eingabelisten enthält.
func Merge(list1, list2 []int) []int {

	result := []int{}

	//Counter initialisieren
	Counter1 := 0
	Counter2 := 0

	//Funktion soll ausgeführt werden, solange beide Counter in Range bleiben
	for Counter1 < len(list1) && Counter2 < len(list2) {

		//Elemente beider Listen vergleichen
		SmallerElement := Min(list1[Counter1], list2[Counter2])

		//Kleineres Elememt an result slice appenden
		result = append(result, SmallerElement)

		if SmallerElement == list1[Counter1] {
			Counter1++
		} else {
			Counter2++
		}
	}

	//Abschlussbedingungen bei ungleich Großen Teillisten
	if Counter1 == len(list1) {
		result = append(result, list2[Counter2:]...)
	} else {
		result = append(result, list1[Counter1:]...)
	}

	return result
}

func Min(a, b int) int {

	if a <= b {
		return a
	}
	return b
}
