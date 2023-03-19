package ordenamientos

// InsertionSort ordena un arreglo de enteros usando el algoritmo de ordenamiento por inserción
func InsertionSort(arreglo []int) {
	for i := 1; i < len(arreglo); i++ {
		j := i
		for j > 0 && arreglo[j-1] > arreglo[j] {
			arreglo[j-1], arreglo[j] = arreglo[j], arreglo[j-1]
			j--
		}
	}
}
