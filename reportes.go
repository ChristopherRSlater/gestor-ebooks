// Paquete reportes: estadísticas generales de la biblioteca.
package reportes

import (
	"fmt"

	"gestor-ebooks/lectura"
)

// TotalLibros devuelve cuántos libros hay registrados.
func TotalLibros(titulos []string) int {
	return len(titulos)
}

// PromedioPaginas devuelve el promedio de páginas por libro.
func PromedioPaginas(paginasTotales []int) float64 {
	if len(paginasTotales) == 0 {
		return 0
	}
	suma := 0
	for i := 0; i < len(paginasTotales); i++ {
		suma += paginasTotales[i]
	}
	return float64(suma) / float64(len(paginasTotales))
}

// LibrosTerminados cuenta los libros leídos por completo.
func LibrosTerminados(paginaActual, paginasTotales []int) int {
	terminados := 0
	for i := 0; i < len(paginasTotales); i++ {
		if lectura.EstaTerminado(paginaActual[i], paginasTotales[i]) {
			terminados++
		}
	}
	return terminados
}

// PosicionLibroMasLargo devuelve la posición del libro con más páginas.
// Si no hay libros, devuelve -1.
func PosicionLibroMasLargo(paginasTotales []int) int {
	if len(paginasTotales) == 0 {
		return -1
	}
	posicion := 0
	for i := 1; i < len(paginasTotales); i++ {
		if paginasTotales[i] > paginasTotales[posicion] {
			posicion = i
		}
	}
	return posicion
}

// MostrarReporte imprime el resumen completo de la biblioteca.
func MostrarReporte(titulos []string, paginasTotales, paginaActual []int) {
	total := TotalLibros(titulos)
	if total == 0 {
		fmt.Println("  (No hay libros para generar un reporte)")
		return
	}

	terminados := LibrosTerminados(paginaActual, paginasTotales)
	enProceso := total - terminados
	posicionLargo := PosicionLibroMasLargo(paginasTotales)

	fmt.Println("  Total de libros       :", total)
	fmt.Println("  Libros terminados     :", terminados)
	fmt.Println("  Libros en proceso     :", enProceso)
	fmt.Printf("  Promedio de páginas   : %.1f\n", PromedioPaginas(paginasTotales))
	fmt.Printf("  Libro más largo       : %s (%d páginas)\n", titulos[posicionLargo], paginasTotales[posicionLargo])
}
