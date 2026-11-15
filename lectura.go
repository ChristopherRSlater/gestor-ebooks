// Paquete lectura: controla el avance de lectura de cada libro.
package lectura

// ActualizarPagina cambia la página actual de UN libro.
// Recibe un APUNTADOR a esa página (*int) para modificar el valor original.
// Devuelve false (y no cambia nada) si la página nueva no es válida.
func ActualizarPagina(paginaActual *int, paginasTotales int, paginaNueva int) bool {
	if paginaNueva < 0 || paginaNueva > paginasTotales {
		return false
	}
	*paginaActual = paginaNueva
	return true
}

// MarcarTerminado coloca la página actual en la última página del libro.
func MarcarTerminado(paginaActual *int, paginasTotales int) {
	*paginaActual = paginasTotales
}

// CalcularProgreso devuelve el porcentaje leído (0 a 100).
// Convierte los enteros a float64 para no perder los decimales en la división.
func CalcularProgreso(paginaActual, paginasTotales int) float64 {
	if paginasTotales == 0 {
		return 0
	}
	return float64(paginaActual) / float64(paginasTotales) * 100
}

// EstaTerminado indica si ya se leyeron todas las páginas.
func EstaTerminado(paginaActual, paginasTotales int) bool {
	return paginaActual >= paginasTotales
}
