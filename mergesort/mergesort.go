package mergesort

// MergeSort sortiert eine Liste von Zahlen mit dem Mergesort-Algorithmus.
func MergeSort(list []int) {

	//Abbruchbedingung
	if len(list) < 2 {
		return
	}

	//Liste wird in zwei teile geteilt
	mid := len(list) / 2
	left := list[:mid]
	right := list[mid:]

	//Teilprozess wiederholen, solange Abbruchbedingung nicht erfüllt ist
	MergeSort(left)
	MergeSort(right)

	//Wird erst ausgeführt, wenn left und right sortiert sind
	//Wenn man direkt list = Merge(left, right) schreiben würde, ändert man nur die Lokale Variable list und nicht die einen Scope höher
	sortedList := Merge(left, right)

	//Um Sicher zu gehen, dass list tatsächlich geändert wird, muss man dieses Copy-Statement hinzufügen
	copy(list, sortedList)
}

/*
Wenn man ein Slice in eine Funktion gibt, dann wird nur der Slice-Header in die Funktion hineingegeben.
Das bedeutet die Funktion erhält einen Pointer zum zugrundeliegenden Array. Die Länge des Slices und die Kapazität des Arrays.
Wenn nun in der Funktion das Array überschrieben werden soll, so kann man das nicht einfach über zuweisungen am Slice machen.

MySlice = {1, 2, 4}

Dieser Code würde keine Veränderung im zugrundeliegenden Array hervorrufen.
Um solch eine Änderung zu erzwingen, muss man die Funktion dazu bringen über den Pointer das zugrundeliegende Array zu accessen.
Dies kann durch Elementweise Zuweisungen oder durch bsp. das Copy() Statement funktionieren.const
*/
