# Sintaxis Tipado y Operadores

Go es un lenguaje **estáticamente tipado**, lo que significa que el compilador sabe el tipo de cada variable. Sin embargo, su sintaxis es limpia y evita la verbosidad de lenguajes como Java.

---

## 1. Declaración de Variables

En Go existen tres formas principales de declarar variables:

### A. Forma Estándar (con `var`)

Se usa principalmente para variables globales o cuando no vas a asignar un valor inmediatamente (inicializa con el **Zero Value**).


```go
var nombre string = "Miguel"
var edad int            // Se inicializa en 0
```

### B. Declaración Corta (`:=`)

Es la más usada dentro de funciones. Go **infiere** el tipo automáticamente.

```go
pais := "Venezuela"     // Go sabe que es string
poblacion := 30000000   // Go sabe que es int
```

### C. Constantes (`const`)

Valores que no cambian en el tiempo.

```go
const Pi = 3.14159
```

---

## 2. Tipos de Datos Nativos

| **Categoría** | **Tipos**                      | **Notas**                                         |
| ------------- | ------------------------------ | ------------------------------------------------- |
| **Enteros**   | `int`, `int8`, `int64`, `uint` | `int` se ajusta a la arquitectura (32 o 64 bits). |
| **Flotantes** | `float32`, `float64`           | `float64` es el estándar por precisión.           |
| **Booleanos** | `bool`                         | `true` o `false`.                                 |
| **Cadenas**   | `string`                       | Inmutables y codificadas en UTF-8.                |

---

## 3. Zero Values (Valores por Defecto)

Si declaras una variable y no le das valor, Go le asigna uno automáticamente. Esto evita errores de "puntero nulo" inesperados.

- **int / float:** `0`
    
- **string:** `""` (string vacío)
    
- **bool:** `false`
    
- **Punteros, Slices, Maps, Interfaces:** `nil`
    

---

## 4. Conversión de Tipos (Casting)

Go **nunca** hace conversión automática (ni siquiera de `int` a `float64`). Debes hacerlo explícitamente.

```go
var a int = 10
var b float64 = float64(a) // Correcto
// var c float64 = a       // Error de compilación
```

---

## 💡 Notas para el Backend

- **CamellCase:** En Go, si una variable o función empieza con **Mayúscula**, es pública (exportada). Si empieza con **minúscula**, es privada del paquete.
    
- **Tipado fuerte:** La rigidez de Go con los tipos es lo que lo hace tan seguro para sistemas de alto rendimiento; los errores se ven al compilar, no al ejecutar.

## 5. Operadores en Go

### A. Aritméticos

Se usan para cálculos matemáticos básicos.

- `+` (Suma)
    
- `-` (Resta)
    
- `*` (Multiplicación)
    
- `/` (División)
    
- `%` (Módulo/Resto)
    

### B. Comparación

Devuelven un valor booleano (`true` o `false`).

- `==` (Igual a)
    
- `!=` (Diferente de)
    
- `<` , `>` (Menor que, Mayor que)
    
- `<=` , `>=` (Menor o igual, Mayor o igual)
    

### C. Lógicos

Para combinar condiciones.

- `&&` (AND lógico)
    
- `||` (OR lógico)
    
- `!` (NOT lógico / Negación)
    

### D. Incremento y Decremento

**¡Cuidado aquí!** A diferencia de otros lenguajes, en Go `++` y `--` son **sentencias**, no expresiones. Esto significa que no puedes hacer `y = x++`. Solo puedes usarlos solos.

- `i++` (Equivale a `i = i + 1`)
    
- `i--` (Equivale a `i = i - 1`)
    

### E. Operadores de Asignación Compuesta

- `+=` (Suma y asigna: `x += 5`)
    
- `-=` (Resta y asigna)
    
- `*=` , `/=` (Multiplica/Divide y asigna)
    

---

## 6. Precedencia de Operadores (Orden de ejecución)

Go sigue el orden matemático estándar, pero aquí tienes una tabla rápida de mayor a menor importancia:

1. `*` `/` `%` `<<` `>>` `&` `&^`
    
2. `+` `-` `|` `^`
    
3. `==` `!=` `<` `<=` `>` `>=`
    
4. `&&`
    
5. `||`
    

> 💡 **Tip de Pro:** Si tienes una expresión compleja, **usa paréntesis `()`**. No solo aseguras el orden, sino que haces que tu código sea mucho más fácil de leer para otros programadores.

---

### 🧠 Resumen

- **Variables:** `nombre := "Go"` (Inferencia).
    
- **Tipos:** Estrictos (no mezcla `int` con `float64`).
    
- **Punteros:** `&` para dirección, `*` para valor.
    
- **Operadores:** `++` solo se usa al final de la variable, nunca antes.
