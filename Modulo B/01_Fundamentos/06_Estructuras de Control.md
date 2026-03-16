Las estructuras de control permiten dirigir el flujo de ejecución. En Go, la sintaxis es limpia: **no se usan paréntesis** en las condiciones, pero las llaves `{ }` son obligatorias.

---

## 1. Condicionales (Toma de Decisiones)

### A. Sentencia `if / else`

Go permite una sintaxis especial llamada **"if con sentencia corta"**, donde declaras una variable que solo existe dentro del bloque del `if`.

```go
// Sintaxis estándar
if edad >= 18 {
    fmt.Println("Mayor de edad")
} else {
    fmt.Println("Menor de edad")
}

// Sentencia corta (Muy usada en manejo de errores)
if valor := calcular(); valor > 100 {
    fmt.Println("Valor alto:", valor)
}
// 'valor' no existe aquí afuera, liberando memoria automáticamente
```

### B. El potente `switch`

En Go, el `switch` es más flexible que en Java. No necesitas `break` (es automático) y puedes evaluar condiciones.

```go
// Switch con valor fijo
dia := "Lunes"
switch dia {
case "Lunes", "Martes", "Miercoles":
    fmt.Println("Día de trabajo")
case "Sabado", "Domingo":
    fmt.Println("Descanso")
default:
    fmt.Println("Día no válido")
}

// Switch sin condición (funciona como un if/else largo más limpio)
numero := 15
switch {
case numero < 0:
    fmt.Println("Negativo")
case numero > 0 && numero < 100:
    fmt.Println("Positivo pequeño")
default:
    fmt.Println("Número grande")
}
```

---

## 2. Bucles e Iteración (El Bucle Único)

En Go **solo existe el `for`**. No hay `while` ni `do-while`, pero el `for` se adapta para cumplir esas funciones.

### A. For clásico (Tres componentes)

```go
for i := 0; i < 5; i++ {
    fmt.Println("Iteración:", i)
}
```

### B. For como "While" (Solo condición)

```go
contador := 0
for contador < 5 {
    fmt.Println(contador)
    contador++
}
```

### C. Bucle Infinito

Ideal para servidores que deben estar siempre escuchando peticiones.

```go
for {
    // Código que corre por siempre
    if detener {
        break // Única forma de salir
    }
}
```

### D. For Range (Iterar colecciones)

Es la forma de oro para recorrer Slices y Maps.

```go
nombres := []string{"Miguel", "Ana", "Jose"}

// Retorna índice y valor
for i, nombre := range nombres {
    fmt.Printf("%d - %s\n", i, nombre)
}

// Si no necesitas el índice, usa el blank identifier (_)
for _, nombre := range nombres {
    fmt.Println(nombre)
}
```

---

## 💡 Tips de Estilo Go

- **Evita el "Else" innecesario:** Si un `if` termina en un `return`, no pongas un `else`. Mantén el código "plano" hacia la izquierda.
    
- **Switch sobre If:** Si tienes más de 3 condiciones sobre la misma variable, usa un `switch`; es mucho más legible.