package quicksort

// QuickSort sortiert die gegebene Liste mit dem Quicksort-Algorithmus.
func QuickSort(list []int) {

	//Abbruchbedingung
	if len(list) < 2 {
		return
	}

	//Der Pivot ist das Letzte Element
	pivotIndex := len(list) - 1
	pivot := list[pivotIndex]

	//Counter i startet bei 0
	i := 0

	//j ranged
	for j := range list {

		//Wenn list[j] kleiner als der Pivot ist
		if list[j] < pivot {
			//Element bei i mit Element bei j vertauschen
			list[i], list[j] = list[j], list[i]
			i++
		}
	}

	//Pivotelement mit i tauschen
	list[i], list[pivotIndex] = list[pivotIndex], list[i]

	QuickSort(list[:i])
	QuickSort(list[i+1:])
}
