package insertionsort

// InsertionSort sortiert die gegebene Liste mit dem Insertion-Sort-Algorithmus.
func InsertionSort(list []int) {

	//Zwei Zahlen vergleichen
	for i := range list {

		//Fall ausschließen, dass das letzte Element verglichen wird (sonst OUT OF RANGE)
		if i == len(list)-1 {
			break
		}

		if list[i] > list[i+1] {
			MoveLeft(list, i+1)
		}
	}

	//Fall 1: Zahl 1 Kleiner als Zahl 2
	//		=> Nichts Tauschen
	//		=> Nächstes Zahlenpaar prüfen

	//Fall 2: Zahl 1 Größer als Zahl 2
	//		=> Zahl 2 verschieben bis keine größere Zahl links von ihr steht
	//		=> Nächstes Zahlenpaar prüfen
}
