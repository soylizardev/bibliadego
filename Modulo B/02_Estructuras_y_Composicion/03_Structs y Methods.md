# Structs y Methods

En Go no hay clases. Usamos **Structs** para agrupar datos y **Methods** para añadirles comportamiento.

## 1. Definición de un Struct

Un struct es una colección de campos.

```go
type Usuario struct {
    ID    int
    Name  string
    Email string
    Activo bool
}
```

## 2. Instanciación (Crear un objeto)

Instanciar no es más que **crear una variable** basada en el molde (el `struct`).

Hay dos formas de hacerlo y una regla de oro:

- **Por posición (Rápida pero peligrosa):**
- 
    ```go
    u := Usuario{1, "Miguel", "migue@mail.com"} 
    // Tienes que poner los datos en el orden exacto en que están en el struct. 
    // Si mañana cambias el orden en el struct, esto rompe.
    ```
    
- **Por nombre de campo (La recomendada):**

    ```go
    u := Usuario{
        Name:  "Miguel",
        Email: "migue@mail.com",
        ID:    1,
    }
    // No importa el orden, es mucho más legible y si el struct cambia, esto no rompe.
    ```
## 3. Methods (El "Comportamiento")

Los métodos se asocian a un struct mediante un **Receiver** (receptor).

### 3.1 El "Receiver" (Receptor)

En otros lenguajes, las funciones viven dentro de la clase. En Go, las funciones están "sueltas", y tú **las "pegas" a un struct** mediante el **Receiver**.

Imagina que el Receiver es como decirle a la función: _"Tú solo le perteneces a este tipo de dato"_.

### A. Receiver por Valor (Copia)

**Receiver por Valor `(u Usuario)`:** Go saca una **fotocopia** del usuario. La función trabaja con la copia. Si cambias el nombre ahí dentro, el original sigue igual. (Úsalo para leer datos).

```go
func (u Usuario) Presentarse() string {
    return "Hola, soy " + u.Name
}
```

### B. Receiver por Puntero (Referencia) ⚡

**CLAVE PARA EL RENDIMIENTO:** Se usa para modificar el struct original o para evitar copiar structs muy grandes en memoria.

```go
func (u *Usuario) Desactivar() {
    u.Activo = false // Modifica el original directamente
}
```

---

## 💡 Mindset de Backend: Composición sobre Herencia

En Go no existe `extends`. Si quieres que un `Admin` tenga los campos de `Usuario`, usamos **Embedding** (Incrustación).

```go
type Admin struct {
    Usuario // El Admin ahora "tiene" ID, Name, Email, etc.
    Nivel   int
}
```

### Embedding (Incrustación / "Herencia" a lo Go)

En Java, dirías: `class Admin extends Usuario`. En Go, decimos: _"El Admin tiene un Usuario adentro"_.

El **Embedding** es poner un struct dentro de otro **sin darle un nombre de campo**. Al hacerlo, el struct de afuera "hereda" automáticamente todos los campos y métodos del de adentro.

```go
type Usuario struct {
    Nombre string
}

type Admin struct {
    Usuario // <--- Esto es EMBEDDING (no tiene nombre de campo)
    Nivel   int
}

// Ahora puedes hacer esto:
var a Admin
a.Nombre = "Miguel" // Accedes directo, como si Nombre fuera del Admin
```

**¿Por qué es mejor que la herencia?** Porque es más simple. No hay jerarquías complicadas de "padres e hijos"; simplemente compones piezas pequeñas para armar una grande.

*Ver [[05_Embedding vs Herencia|Embedding vs Herencia]]* para profundizar más en el tema.

### Comparativa Java vs Go

| **Característica** | **Java**                   | **Go**                                     |
| ------------------ | -------------------------- | ------------------------------------------ |
| **Contenedor**     | `class`                    | `struct`                                   |
| **Comportamiento** | Métodos dentro de la clase | Métodos fuera, unidos por un **Receiver**  |
| **Herencia**       | `extends`                  | **Embedding** (Incrustación de structs)    |
| **Polimorfismo**   | Herencia de clases         | **Interfaces** (Basadas en comportamiento) |

![](https://youtu.be/c8H0w4yBL10)

### Ejercicio:

### 🛠️ Reto: Gestión de Productos

**Objetivo:** Crear un programa que gestione el precio de un producto usando métodos con diferentes tipos de receivers.

#### 1. Definir la Estructura (El Molde)

Crea un struct llamado `Producto` que tenga:

- `Nombre` (string)
    
- `Precio` (float64)
    

#### 2. Crear los Métodos (El Comportamiento)

- **Método `Mostrar()`:** Debe tener un **Receiver por Valor**. Su única función es imprimir en consola: ```"Producto: [Nombre] - Precio: $[Precio]".
    
- **Método `AplicarDescuento()`:** Debe recibir un `porcentaje` (float64) y tener un **Receiver por Puntero**. Debe calcular el nuevo precio y actualizarlo en el struct original.
    

#### 3. Ejecución en el `main`

- **Instancia** un producto usando la **forma explícita** (nombres de campos). Ejemplo: `p := Producto{Nombre: "Laptop", Precio: 1000}`.
    
- Llama al método `Mostrar()` para ver el precio inicial.
    
- Llama al método `AplicarDescuento(10)` (para un 10% de descuento).
    
- Llama de nuevo a `Mostrar()` para comprobar que el precio bajó de verdad.
    

---

### 💡 Pista para el éxito:

Si el método `AplicarDescuento` no tuviera el asterisco `*` en el receiver, cuando vuelvas a llamar a `Mostrar()`, verías el precio viejo. El `*` es la llave que le permite a la función entrar a la memoria y cambiar el dato original.

### Respuesta

```go
package main

import "fmt"

type producto struct {
	nombre string
	precio float64
}

func (p producto) mostrar() {
	fmt.Printf("Producto: %s - Precio: $%f\n", p.nombre, p.precio)
}

func (p *producto) AplicarDescuento(porcentaje float64) {
	p.precio= p.precio - p.precio * (porcentaje /100)
}

func main()  {
	
	p := producto {
		nombre: "laptop",
		precio: 1000,
	}

	p.mostrar()
	p.AplicarDescuento(10)
	p.mostrar()

}
``` 
