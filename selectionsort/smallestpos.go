package selectionsort

// SmallestPos gibt die Position des kleinsten Elements in der Liste zurück.
func SmallestPos(list []int) int {
	smallest := 0

	for i := range list {

		if list[i] < list[smallest] {

			smallest = i
		}

	}
	return smallest
}
