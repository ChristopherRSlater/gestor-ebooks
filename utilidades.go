// Paquete utilidades: entrada de datos y validaciones reutilizables.
//
// Todas las funciones que empiezan con mayúscula son "exportadas",es decir, se pueden usar desde otros paquetes (por ejemplo desde main).
package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// lector es una variable de paquete: lee lo que el usuario escribe en consola. **
var lector = bufio.NewReader(os.Stdin)

// leerLinea le una línea completa del teclado y le quita espacios y saltos de línea.
// Si la entrada se cierra (Ctrl+Z / Ctrl+D), el programa termina de forma ordenada.
func leerLinea() string {
	linea, err := lector.ReadString('\n')
	if err != nil && len(linea) == 0 {
		fmt.Println("\nEntrada finalizada. Hasta pronto.")
		os.Exit(0)
	}
	return strings.TrimSpace(linea)
}

// LeerTexto muestra un mensaje y pide un texto que NO puede estar vacío.
func LeerTexto(mensaje string) string {
	for {
		fmt.Print(mensaje)
		texto := leerLinea()
		if texto != "" {
			return texto
		}
		fmt.Println("  [!] El texto no puede estar vacío. Intenta de nuevo.")
	}
}

// LeerEntero muestra un mensaje y pide un número entero.
// Convierte el texto a int con strconv.Atoi y repite hasta que sea válido.
func LeerEntero(mensaje string) int {
	for {
		fmt.Print(mensaje)
		texto := leerLinea()
		numero, err := strconv.Atoi(texto)
		if err != nil {
			fmt.Println("  [!] Debes escribir un número entero. Intenta de nuevo.")
			continue
		}
		return numero
	}
}

// LeerEnteroEnRango pide un entero y valida que esté entre minimo y maximo (inclusive).
func LeerEnteroEnRango(mensaje string, minimo, maximo int) int {
	for {
		numero := LeerEntero(mensaje)
		if numero >= minimo && numero <= maximo {
			return numero
		}
		fmt.Printf("  [!] El número debe estar entre %d y %d.\n", minimo, maximo)
	}
}

// Confirmar hace una pregunta de sí/no y devuelve true si la respuesta es "s" o "si".
func Confirmar(mensaje string) bool {
	fmt.Print(mensaje + " (s/n): ")
	respuesta := strings.ToLower(leerLinea())
	return respuesta == "s" || respuesta == "si" || respuesta == "sí"
}

// Pausar espera a que el usuario presione Enter.
func Pausar() {
	fmt.Print("\nPresiona Enter para continuar...")
	leerLinea()
}

// Separador imprime una línea decorativa para ordenar la salida en pantalla.
func Separador() {
	fmt.Println(strings.Repeat("=", 60))
}
