package main

import (
	"busquedas"
	"fmt"
	"sort"
)

func main() {
	arreglo := []int{-1, 5, -6, 10, 7, 2, -10, 5, 3, 17}
	buscado := 3
	// Busqueda Lineal
	fmt.Println(busquedas.BusLineal(arreglo, buscado))

	// Ordenar el arreglo
	sort.Ints(arreglo)
	fmt.Println(arreglo)
	// Busqueda Binaria
	fmt.Println(busquedas.BusquedaBinaria(arreglo, 5))
}
