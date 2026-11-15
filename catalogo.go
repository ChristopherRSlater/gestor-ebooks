// Paquete catalogo: administra la lista de libros electrónicos.--------------
//
// Cada libro vive en la MISMA posición (índice) de cuatro slices paralelos:
//
//	titulos[i], autores[i], paginasTotales[i], paginaActual[i]
//
// Las funciones que agregan o eliminan libros reciben APUNTADORES a los slices
//  (*[]string, *[]int) para poder modificar los slices originales de main.
package catalogo

import (
	"fmt"
	"strings"

	"gestor-ebooks/lectura"
)

// AgregarLibro añade un libro nuevo al final de los cuatro slices.
// Un libro nuevo empieza en la página 0 (sin leer).
func AgregarLibro(titulos, autores *[]string, paginasTotales, paginaActual *[]int,
	titulo, autor string, paginas int) {

	*titulos = append(*titulos, titulo)
	*autores = append(*autores, autor)
	*paginasTotales = append(*paginasTotales, paginas)
	*paginaActual = append(*paginaActual, 0)
}

// ListarLibros imprime todos los libros en forma de tabla.
func ListarLibros(titulos, autores []string, paginasTotales, paginaActual []int) {
	if len(titulos) == 0 {
		fmt.Println("  (No hay libros registrados todavía)")
		return
	}

	fmt.Printf("%-4s %-28s %-20s %-8s %-10s\n", "N°", "Título", "Autor", "Páginas", "Progreso")
	fmt.Println(strings.Repeat("-", 76))

	for i := 0; i < len(titulos); i++ {
		porcentaje := lectura.CalcularProgreso(paginaActual[i], paginasTotales[i])
		estado := fmt.Sprintf("%.0f%%", porcentaje)
		if lectura.EstaTerminado(paginaActual[i], paginasTotales[i]) {
			estado = "Terminado"
		}
		// %-28.27s corta el texto a 27 caracteres para que quepa en la columna.
		fmt.Printf("%-4d %-28.27s %-20.19s %-8d %-10s\n",
			i+1, titulos[i], autores[i], paginasTotales[i], estado)
	}
}

// BuscarPorTitulo imprime los libros cuyo título contiene el texto buscado
// (sin distinguir mayúsculas) y devuelve cuántos encontró.
func BuscarPorTitulo(titulos, autores []string, paginasTotales, paginaActual []int, texto string) int {
	encontrados := 0
	buscado := strings.ToLower(texto)

	for i := 0; i < len(titulos); i++ {
		if strings.Contains(strings.ToLower(titulos[i]), buscado) {
			encontrados++
			fmt.Printf("  %d) %s — %s (%d/%d páginas)\n",
				i+1, titulos[i], autores[i], paginaActual[i], paginasTotales[i])
		}
	}
	return encontrados
}

// EliminarLibro quita el libro de la posición indice en los cuatro slices.
// Devuelve false si el índice no es válido.
func EliminarLibro(titulos, autores *[]string, paginasTotales, paginaActual *[]int, indice int) bool {
	if indice < 0 || indice >= len(*titulos) {
		return false
	}

	quitarTexto(titulos, indice)
	quitarTexto(autores, indice)
	quitarEntero(paginasTotales, indice)
	quitarEntero(paginaActual, indice)
	return true
}

// quitarTexto reconstruye la lista de textos sin el elemento de la posición indice.
// Recibe un APUNTADOR para reemplazar la lista original por la nueva.
func quitarTexto(lista *[]string, indice int) {
	nueva := []string{}
	for i := 0; i < len(*lista); i++ {
		if i != indice {
			nueva = append(nueva, (*lista)[i])
		}
	}
	*lista = nueva
}

// quitarEntero hace lo mismo que quitarTexto, pero con una lista de enteros.
func quitarEntero(lista *[]int, indice int) {
	nueva := []int{}
	for i := 0; i < len(*lista); i++ {
		if i != indice {
			nueva = append(nueva, (*lista)[i])
		}
	}
	*lista = nueva
}
