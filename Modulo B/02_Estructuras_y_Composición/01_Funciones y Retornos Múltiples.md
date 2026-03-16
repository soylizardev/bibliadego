En Go, las funciones son "ciudadanos de primera clase" (_First-class citizens_). Esto significa que pueden ser asignadas a variables, pasadas como argumentos y retornadas por otras funciones.

## 1. Anatomía de una Función

A diferencia de Java, el tipo de dato va **después** del nombre del parámetro.

```go
func sumar(a int, b int) int {
    return a + b
}
```

## 2. El Superpoder: Retornos Múltiples

Esta es la característica más icónica de Go. Se usa casi siempre para devolver un **resultado** y un **error** al mismo tiempo.

```go
func dividir(dividendo, divisor float64) (float64, error) {
    if divisor == 0 {
        return 0, fmt.Errorf("no puedes dividir por cero")
    }
    return dividendo / divisor, nil
}

func main() {
    res, err := dividir(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Resultado:", res)
}
```

> [!IMPORTANT] En Go, no usamos `try-catch`. Usamos retornos múltiples para obligar al programador a lidiar con el error en el momento.

## 3. Retornos Nombrados (Named Returns)

Puedes pre-declarar las variables que vas a retornar en la firma de la función. Esto hace que el código sea más legible en funciones cortas.

Go

```go
func rectangulo(largo, ancho float64) (area, perimetro float64) {
    area = largo * ancho
    perimetro = (largo + ancho) * 2
    return // Retorna automáticamente area y perimetro
}
```

## 4. Funciones Variádicas

Son funciones que aceptan un número indefinido de argumentos (como `fmt.Println`). Se usa el prefijo `...`

```go
func sumarTodo(numeros ...int) int {
    total := 0
    for _, n := range numeros {
        total += n
    }
    return total
}
```

## 5. Funciones Anónimas y Closures

Puedes declarar una función sin nombre dentro de otra y ejecutarla al instante o guardarla en una variable.

```go
func main() {
    saludo := func(nombre string) {
        fmt.Printf("Hola %s desde una función anónima\n", nombre)
    }
    saludo("Miguel")
}
```

---

### 💡 Conexión con las Interfaces[^1]

- Una interfaz define **firmas de funciones**.
    
- Si una función de interfaz dice que devuelve `(string, error)`, tu struct debe tener un método que devuelva **exactamente eso**.
---
[^1]: Ver [[04_Interfaces|Interfaces]] para mayor información
