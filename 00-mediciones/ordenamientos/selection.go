package ordenamientos

// SelectionSort ordena un arreglo de enteros usando el algoritmo de ordenamiento por selección
func SelectionSort(arreglo []int) {
	for i := 0; i < len(arreglo); i++ {
		min := i
		for j := i + 1; j < len(arreglo); j++ {
			if arreglo[j] < arreglo[min] {
				min = j
			}
		}
		arreglo[i], arreglo[min] = arreglo[min], arreglo[i]
	}
}
