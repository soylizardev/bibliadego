# Generics

Los **Generics** nos permiten escribir funciones y estructuras de datos que pueden trabajar con múltiples tipos sin sacrificar la seguridad del tipado. Evitan la duplicación de código ("Don't Repeat Yourself" - DRY).

---

## 1. El Problema: Duplicación de Código

Antes de los genéricos, si querías una función para sumar, necesitabas una para cada tipo:

```go
func sumarEnteros(a, b int) int { return a + b }
func sumarFloats(a, b float64) float64 { return a + b }
```

_Esto es ineficiente porque la lógica es la misma, solo cambia el tipo._

---

## 2. La Solución: Funciones Genéricas

Con genéricos, declaras un **parámetro de tipo** (por convención se usa `T`) entre corchetes `[]` antes de los parámetros de la función.

### Ejemplo de Función Sumar:

```go
// T puede ser int o float64
func Sumar[T int | float64](a, b T) T {
    return a + b
}

func main() {
    // La misma función sirve para ambos casos
    fmt.Println(Sumar(10, 20))       // Funciona con int
    fmt.Println(Sumar(1.5, 2.5))     // Funciona con float64
}
```

---

## 3. Elementos Clave de la Sintaxis

- **Corchetes `[]`**: Aquí es donde defines las restricciones del genérico.
    
- **Tipo `T`**: Es el nombre convencional del tipo genérico (puedes ponerle cualquier nombre, pero `T` es el estándar en Go).
    
- **Constraints (Restricciones)**: El símbolo `|` (pipe) actúa como un "O". En el ejemplo: `int | float64` significa que la función acepta cualquiera de los dos.
    

---

## 💡 Por qué es "Alto Rendimiento"

- **Seguridad en tiempo de compilación**: A diferencia de usar `interface{}` (que veremos luego), los genéricos no tienen penalización de rendimiento en ejecución porque el compilador genera el código específico para cada tipo que uses.
    
- **Código Limpio**: Reduces drásticamente la cantidad de líneas de código en proyectos grandes.
