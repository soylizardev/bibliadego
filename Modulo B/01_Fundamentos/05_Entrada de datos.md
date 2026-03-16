# Entrada de datos (Scan)

En Go, pedir datos al usuario requiere entender cómo se comporta el **Buffer** (la memoria temporal donde se guarda lo que tecleas).

## 1. La Familia `fmt` (Lectura básica)

### A. `fmt.Scan`

Lee valores separados por espacios o saltos de línea. Ideal para números o palabras sueltas.

```go
var edad int
fmt.Print("Ingresa tu edad: ")
fmt.Scan(&edad) // Importante el puntero &
```

### B. `fmt.Scanln`

Similar a `Scan`, pero deja de leer en cuanto encuentra un **salto de línea** (Enter). Útil para leer varios datos en una sola línea.

```go
var dia, mes int
fmt.Scanln(&dia, &mes) // Ejemplo: 15 03
```

### C. `fmt.Scanf` (Lectura con Formato)

Permite obligar al usuario a escribir con un formato específico (ej. una hora con dos puntos).

```go
var hora, minuto int
fmt.Scanf("%d:%d", &hora, &minuto) // El usuario DEBE escribir "13:21"
```

---

## 2. Lectura Avanzada con `bufio` (Frases completas)

El problema de `fmt.Scan` es que si escribes _"Hola Mundo"_, solo captura _"Hola"_. Para leer frases con espacios, usamos `bufio` y el paquete `os`.

```go
import (
    "bufio"
    "os"
    "strings"
)

// 1. Creamos un "Lector" que escucha el teclado (Stdin)
reader := bufio.NewReader(os.Stdin)

fmt.Print("Escribe una frase: ")
// 2. Lee hasta que encuentra el salto de línea (\n)
frase, _ := reader.ReadString('\n') 

// 3. Limpieza: ReadString incluye el "Enter" al final. 
// Usamos strings.TrimSpace para quitar ese salto de línea
fraseLimpia := strings.TrimSpace(frase)
```

---

## ⚠️ Concepto Crítico: Limpiar el Buffer

Cuando mezclas `Scan` con `bufio`, a veces el programa "se salta" una pregunta. Esto pasa porque `Scan` deja el "Enter" (`\n`) flotando en la memoria y el siguiente lector lo captura pensando que ya terminaste.

**Solución:** Crear un lector para "consumir" esos saltos de línea sobrantes antes de pedir una frase larga.

---

## 💡 Resumen de uso:

- **¿Solo un número o palabra?** -> `fmt.Scan`.
    
- **¿Varios datos en una línea?** -> `fmt.Scanln`.
    
- **¿Una frase larga (nombre completo, descripción)?** -> `bufio.NewReader`.
