# Entrada de datos (Scan y Bufio)

En Go, pedir datos al usuario requiere entender cómo se comporta el **Buffer** (la memoria temporal donde se guarda lo que tecleas) y cómo el lenguaje accede a las variables mediante punteros.

## 1. La Familia `fmt` (Lectura básica)

### A. `fmt.Scan`

Lee valores separados por espacios o saltos de línea. Es la forma más común de capturar números o palabras sueltas.

```go
var edad int
fmt.Print("Ingresa tu edad: ")
// Se usa el símbolo & para que Scan sepa en qué dirección de memoria guardar el dato
fmt.Scan(&edad) 
```

### B. `fmt.Scanln`

Similar a `Scan`, pero termina la lectura inmediatamente al encontrar un **salto de línea** (cuando el usuario pulsa Enter). Es ideal para capturar varios datos que el usuario ingresa en una sola fila.

```go
var dia, mes int
fmt.Print("Ingresa día y mes (ej: 15 03): ")
fmt.Scanln(&dia, &mes)
```

### C. `fmt.Scanf` (Lectura con Formato)

Permite extraer datos que siguen un patrón o formato específico definido por nosotros, como si fuera una plantilla.

```go
var hora, minuto int
fmt.Print("Ingresa la hora (HH:MM): ")
// El usuario debe escribir los dos puntos para que los datos se capturen correctamente
fmt.Scanf("%d:%d", &hora, &minuto)
```

---

## 2. Lectura Avanzada con `bufio` (Frases completas)

El paquete `fmt` deja de leer cuando encuentra el primer espacio. Si quieres capturar una frase completa (como un nombre con apellido), necesitas usar el paquete `bufio` junto con `os`.

```go
import (
    "bufio"
    "os"
    "strings"
)

// 1. Creamos un "Lector" que escucha la entrada estándar (teclado)
reader := bufio.NewReader(os.Stdin)

fmt.Print("Ingresa tu nombre completo: ")

// 2. Lee todo el texto incluyendo espacios hasta presionar Enter (\n)
texto, _ := reader.ReadString('\n')

// 3. Limpieza: ReadString guarda el texto con el "Enter" al final.
// Usamos strings.TrimSpace para quitar ese salto de línea y espacios extra
textoLimpio := strings.TrimSpace(texto)
```

---

## 3. El Problema del Buffer (Salto de Línea atrapado)

Un error común al mezclar `fmt.Scan` con `bufio` es que el programa parece "saltarse" una pregunta.

### ¿Por qué sucede?

Cuando usas `fmt.Scan`, el usuario escribe el dato y pulsa **Enter**. El dato se guarda en la variable, pero el **salto de línea (`\n`)** del Enter se queda flotando en el buffer. Si la siguiente instrucción es un `ReadString`, este verá ese "Enter" residual, pensará que ya terminaste de escribir y continuará sin dejarte teclear nada.

### Solución: Limpiar el Buffer

Para evitar esto, después de usar un `Scan` y antes de un `bufio`, debemos limpiar el buffer leyendo ese salto de línea sobrante:

```go
// Después de un fmt.Scan(&variable)
reader := bufio.NewReader(os.Stdin)
reader.ReadString('\n') // Esto "se come" el Enter sobrante y deja el buffer limpio
```

---

## 💡 Resumen de uso:

- **Para una sola palabra o número:** Usa `fmt.Scan`.
    
- **Para datos con formato (ej. fechas 12/10):** Usa `fmt.Scanf`.
    
- **Para frases con espacios:** Usa `bufio.NewReader`.
    
- **Punteros:** No olvides nunca el símbolo `&` antes de la variable en la familia `Scan`.