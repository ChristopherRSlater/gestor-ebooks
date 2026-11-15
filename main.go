// Sistema de Gestión de Libros Electrónicos (e-books)
// Autónomo 1 — Programación en Go (Unidad 1)
//
// Programa de consola organizado con programación funcional:
// funciones pequeñas agrupadas en paquetes, sin structs.
package main

import (
	"fmt"

	"gestor-ebooks/catalogo"
	"gestor-ebooks/lectura"
	"gestor-ebooks/reportes"
	"gestor-ebooks/utilidades"
)

func main() {
	// Datos del sistema: cuatro slices paralelos (el libro i está en la posición i).
	// Se cargan 3 libros de ejemplo para poder probar el programa rápidamente.
	titulos := []string{"Cien años de soledad", "El principito", "Don Quijote de la Mancha"}
	autores := []string{"Gabriel García Márquez", "Antoine de Saint-Exupéry", "Miguel de Cervantes"}
	paginasTotales := []int{471, 96, 1072}
	paginaActual := []int{120, 96, 0}

	opcion := -1
	for opcion != 0 {
		mostrarMenu()
		opcion = utilidades.LeerEnteroEnRango("Elige una opción: ", 0, 7)
		fmt.Println()

		switch opcion {
		case 1:
			opcionAgregar(&titulos, &autores, &paginasTotales, &paginaActual)
		case 2:
			catalogo.ListarLibros(titulos, autores, paginasTotales, paginaActual)
		case 3:
			opcionBuscar(titulos, autores, paginasTotales, paginaActual)
		case 4:
			opcionEliminar(&titulos, &autores, &paginasTotales, &paginaActual)
		case 5:
			opcionActualizarPagina(titulos, autores, paginasTotales, paginaActual)
		case 6:
			opcionMarcarTerminado(titulos, autores, paginasTotales, paginaActual)
		case 7:
			reportes.MostrarReporte(titulos, paginasTotales, paginaActual)
		case 0:
			fmt.Println("¡Gracias por usar el gestor de e-books! Hasta pronto.")
		}

		if opcion != 0 {
			utilidades.Pausar()
		}
	}
}

// mostrarMenu imprime las opciones disponibles.
func mostrarMenu() {
	fmt.Println()
	utilidades.Separador()
	fmt.Println("   GESTOR DE LIBROS ELECTRÓNICOS")
	utilidades.Separador()
	fmt.Println("   1. Agregar libro")
	fmt.Println("   2. Listar libros")
	fmt.Println("   3. Buscar libro por título")
	fmt.Println("   4. Eliminar libro")
	fmt.Println("   5. Actualizar página de lectura")
	fmt.Println("   6. Marcar libro como terminado")
	fmt.Println("   7. Ver reporte general")
	fmt.Println("   0. Salir")
	utilidades.Separador()
}

// opcionAgregar pide los datos de un libro nuevo y lo registra.
func opcionAgregar(titulos, autores *[]string, paginasTotales, paginaActual *[]int) {
	fmt.Println("--- Agregar libro ---")
	titulo := utilidades.LeerTexto("Título: ")
	autor := utilidades.LeerTexto("Autor: ")
	paginas := utilidades.LeerEnteroEnRango("Número de páginas (1-5000): ", 1, 5000)

	catalogo.AgregarLibro(titulos, autores, paginasTotales, paginaActual, titulo, autor, paginas)
	fmt.Println("  ✔ Libro agregado correctamente.")
}

// opcionBuscar pide un texto y muestra los libros que coinciden.
func opcionBuscar(titulos, autores []string, paginasTotales, paginaActual []int) {
	fmt.Println("--- Buscar libro ---")
	texto := utilidades.LeerTexto("Texto a buscar en el título: ")

	encontrados := catalogo.BuscarPorTitulo(titulos, autores, paginasTotales, paginaActual, texto)
	if encontrados == 0 {
		fmt.Println("  No se encontraron libros con ese texto.")
	}
}

// opcionEliminar muestra la lista, pide un número y elimina ese libro.
func opcionEliminar(titulos, autores *[]string, paginasTotales, paginaActual *[]int) {
	fmt.Println("--- Eliminar libro ---")
	indice := elegirLibro(*titulos, *autores, *paginasTotales, *paginaActual)
	if indice < 0 {
		return
	}

	if utilidades.Confirmar("¿Seguro que quieres eliminar \"" + (*titulos)[indice] + "\"?") {
		catalogo.EliminarLibro(titulos, autores, paginasTotales, paginaActual, indice)
		fmt.Println("  ✔ Libro eliminado.")
	} else {
		fmt.Println("  Operación cancelada.")
	}
}

// opcionActualizarPagina cambia la página en la que va el lector.
func opcionActualizarPagina(titulos, autores []string, paginasTotales, paginaActual []int) {
	fmt.Println("--- Actualizar página de lectura ---")
	indice := elegirLibro(titulos, autores, paginasTotales, paginaActual)
	if indice < 0 {
		return
	}

	nueva := utilidades.LeerEntero("Página en la que vas: ")
	// &paginaActual[indice] es un apuntador a la página de ese libro.
	if lectura.ActualizarPagina(&paginaActual[indice], paginasTotales[indice], nueva) {
		porcentaje := lectura.CalcularProgreso(paginaActual[indice], paginasTotales[indice])
		fmt.Printf("  ✔ Progreso actualizado: %.1f%%\n", porcentaje)
	} else {
		fmt.Printf("  [!] Página inválida. Debe estar entre 0 y %d.\n", paginasTotales[indice])
	}
}

// opcionMarcarTerminado marca un libro como leído por completo.
func opcionMarcarTerminado(titulos, autores []string, paginasTotales, paginaActual []int) {
	fmt.Println("--- Marcar libro como terminado ---")
	indice := elegirLibro(titulos, autores, paginasTotales, paginaActual)
	if indice < 0 {
		return
	}

	lectura.MarcarTerminado(&paginaActual[indice], paginasTotales[indice])
	fmt.Println("  ✔ ¡Libro terminado! Felicitaciones por la lectura.")
}

// elegirLibro lista los libros y pide al usuario el número de uno.
// Devuelve su posición (índice) o -1 si no hay libros registrados.
func elegirLibro(titulos, autores []string, paginasTotales, paginaActual []int) int {
	if len(titulos) == 0 {
		fmt.Println("  (No hay libros registrados todavía)")
		return -1
	}

	catalogo.ListarLibros(titulos, autores, paginasTotales, paginaActual)

	numero := utilidades.LeerEnteroEnRango("\nNúmero del libro: ", 1, len(titulos))
	return numero - 1
}
