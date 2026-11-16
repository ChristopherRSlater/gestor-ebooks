# Sistema de Gestión de Libros Electrónicos (e-books)

Proyecto del **Autónomo 1** — Programación en Go (Unidad 1)
Universidad Internacional del Ecuador

## Descripción

Programa de consola escrito en Go que permite administrar una biblioteca personal de libros electrónicos: registrar libros, buscarlos, controlar en qué página va el lector y consultar estadísticas generales.

El software está desarrollado con **programación funcional**: el código se organiza en funciones pequeñas (cada una con una sola tarea) agrupadas en paquetes. 

## Objetivo

Ofrecer una herramienta sencilla para gestionar libros electrónicos y llevar el control del progreso de lectura de cada uno.

## Alcance

**Incluye:** menú interactivo en consola, validación de datos, registro y consulta de libros, control de progreso y reporte general de la lectura.

**No incluye:** interfaz gráfica, base de datos ni almacenamiento en archivos (los datos viven en memoria mientras el programa está abierto), tampoco incluye web scraping. Estas son posibles mejoras futuras.

## Estructura del proyecto

```
gestor-ebooks/
├── go.mod                     # definición del módulo
├── main.go                    # menú principal y flujo del programa
├── catalogo/catalogo.go       # agregar, listar, buscar y eliminar libros
├── lectura/lectura.go         # progreso de lectura
├── reportes/reportes.go       # estadísticas generales
└── utilidades/utilidades.go   # entrada de datos y validaciones
```

## Módulos y funciones

| Paquete | Funciones | Descripción |
|---|---|---|
| `main` | `main`, `mostrarMenu`, `opcion...`, `elegirLibro` | Menú principal con `switch` y ciclo `for` |
| `catalogo` | `AgregarLibro`, `ListarLibros`, `BuscarPorTitulo`, `EliminarLibro` | Gestión de la lista de libros |
| `lectura` | `ActualizarPagina`, `MarcarTerminado`, `CalcularProgreso`, `EstaTerminado` | Avance de lectura |
| `reportes` | `TotalLibros`, `PromedioPaginas`, `LibrosTerminados`, `PosicionLibroMasLargo`, `MostrarReporte` | Estadísticas |
| `utilidades` | `LeerTexto`, `LeerEntero`, `LeerEnteroEnRango`, `Confirmar`, `Pausar`, `Separador` | Entrada de datos y validaciones |


## Paquetes utilizados

Solo se usa la **librería estándar de Go** (no hay paquetes de terceros):

| Paquete | Uso |
|---|---|
| `fmt` | Mostrar textos y tablas en pantalla |
| `bufio` | Leer líneas del teclado |
| `os` | Acceso a la entrada estándar y salida del programa |
| `strings` | Recortar, buscar y repetir textos |
| `strconv` | Convertir texto a número entero |

## Cómo ejecutar

Requisito: tener [Go](https://go.dev/dl/) 1.21 o superior instalado.

```bash
git clone https://github.com/ChristopherRSlater/gestor-ebooks.git
cd gestor-ebooks
go run .
```

Para generar el ejecutable:

```bash
go build -o gestor-ebooks.exe .
```

## Autor

Christopher Slater— Universidad Internacional del Ecuador
