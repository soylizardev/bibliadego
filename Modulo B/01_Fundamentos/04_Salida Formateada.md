## (`fmt.Printf`)

A diferencia de `Println`, **`Printf`** (Print Formatted) nos permite usar **verbos** (especificadores) para controlar exactamente cómo se ven los datos.

---

## 1. Verbos de Formato Comunes

Estos son los "comodines" que usas dentro del string para representar variables:

| **Verbo** | **Descripción**                           | **Ejemplo de Uso**                  |
| --------- | ----------------------------------------- | ----------------------------------- |
| `%s`      | **Strings** (Cadenas de texto)            | `fmt.Printf("Hola %s", nombre)`     |
| `%d`      | **Integers** (Enteros en base 10)         | `fmt.Printf("Edad: %d", 25)`        |
| `%f`      | **Floats** (Números decimales)            | `fmt.Printf("Altura: %f", 1.80)`    |
| `%.2f`    | **Floats con precisión** (2 decimales)    | `fmt.Printf("%.2f", 3.14159)`       |
| `%t`      | **Booleans** (true o false)               | `fmt.Printf("Activo: %t", true)`    |
| `%v`      | **Default Value** (El valor "tal cual")   | Útil cuando no sabes el tipo exacto |
| `%T`      | **Type** (Imprime el tipo de la variable) | `fmt.Printf("Es un: %T", valor)`    |

---

## 2. Bases Numéricas y Conversiones

Go facilita mucho ver un número en diferentes sistemas desde la consola:

- `%b`: Binario.
    
- `%o`: Octal.
    
- `%x`: Hexadecimal.
    
- `%e`: Notación científica.

```go
numero := 255
fmt.Printf("Binario: %b\n", numero) // 11111111
fmt.Printf("Hex: %x\n", numero)     // ff
```

---

## 3. Alineación y Tablas

Puedes definir el ancho de un campo para alinear texto a la izquierda o derecha, ideal para reportes en consola:

```go
// %-10s -> Alinea string a la izquierda con 10 espacios
// %3d   -> Alinea entero a la derecha con 3 espacios
fmt.Printf("%-10s | %3d\n", "Carlos", 25)
fmt.Printf("%-10s | %3d\n", "Ana", 30)
```

---

## 4. El "Truco" de `fmt.Sprintf` 🌟

Este es uno de los puntos más importantes del video. A veces no quieres imprimir en consola, sino **guardar el texto formateado en una variable** (para enviarlo por email, guardarlo en un log o devolverlo en una API).

```go
// Sprintf NO imprime, DEVUELVE un string
mensaje := fmt.Sprintf("Usuario %s (ID: %d) creado.", "Miguel", 101)

fmt.Println(mensaje) // Ahora sí lo imprimes si quieres
```

---

## 💡 Tips

- **Salto de línea:** `Printf` no añade salto de línea al final. Debes poner siempre `\n` manualmente para que la siguiente impresión no salga pegada.
    
- **Preferencia:** Usa `%.2f` siempre que manejes dinero o medidas para evitar que salgan 6 decimales por defecto.