package ordenamientos

// BubbleSort ordena un arreglo de enteros usando el algoritmo de ordenamiento por burbuja
func BubbleSort(arreglo []int) {
	for i := 0; i < len(arreglo); i++ {
		for j := 0; j < len(arreglo)-1; j++ {
			if arreglo[j] > arreglo[j+1] {
				arreglo[j], arreglo[j+1] = arreglo[j+1], arreglo[j]
			}
		}
	}
}
