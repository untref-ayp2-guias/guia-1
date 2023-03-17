package busquedas

func BusquedaBinaria(datos []int, buscado int) int {
	inicio := 0
	fin := len(datos) - 1

	for inicio <= fin {
		medio := (inicio + fin) / 2
		switch {
		case buscado < datos[medio]:
			fin = medio - 1
		case buscado > datos[medio]:
			inicio = medio + 1
		default:
			return medio
		}
	}
	return -1
}
