# 🏷️ Custom Types (Tipos Personalizados)

Los tipos personalizados en Go permiten crear alias o tipos basados en tipos primitivos existentes (como `string`, `int`, `float64`).

**¿Para qué sirven?**

1. **Claridad:** El código es más legible (no es lo mismo un `string` que un `IDUsuario`).
    
2. **Seguridad:** Evitan errores al restringir valores o tipos de datos específicos.
    
3. **Poder:** Permiten añadir métodos a tipos básicos (algo que veremos más adelante).
    

---

## 1. Definición Básica

Puedes crear un tipo a partir de cualquier tipo primitivo usando la palabra clave `type`.

```go
// Definimos que 'edad' es un tipo basado en 'int'
type edad int

func main() {
    var miEdad edad = 25
    fmt.Println(miEdad) // Imprime: 25
}
```

---

## 2. Uso Práctico con Constantes (Enums a lo Go)

Esto es vital para estados de pedidos, roles de usuario, etc.

### Ejemplo de "Kages":

```go
// 1. Creamos el tipo base
type Kage string

// 2. Definimos las constantes válidas para ese tipo
const (
    Hokage   Kage = "Hokage"
    Kazekage Kage = "Kazekage"
    Mizukage Kage = "Mizukage"
)

func main() {
    // Uso por declaración explícita
    var lider Kage = Hokage
    
    // Uso por inferencia de tipo
    actualLider := Kazekage

    fmt.Println("El líder actual es:", lider) //
}
```

---

## 💡 Ventajas Clave del Video

- **Restricción de valores:** Si intentas asignar un valor que no esté definido en tus constantes (ej. `KageX`), el compilador te ayudará a detectar que algo no está bien.
    
- **Previsibilidad:** Al usar estos tipos en funciones, te aseguras de que nadie pase un `string` genérico cuando esperas un tipo de dato específico del dominio de tu negocio.
    

---
### 🔑Importancia

Este concepto es el puente hacia las **Interfaces**. En Go, a menudo creamos tipos personalizados para luego hacer que "cumplan" con una interfaz.