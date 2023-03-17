package busquedas

func BusLineal(datos []int, buscado int) int {
	for indice, valor := range datos {
		if buscado == valor {
			return indice
		}
	}
	return -1
}
