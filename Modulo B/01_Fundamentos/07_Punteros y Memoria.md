# Punteros y Memoria

### 1. ¿Qué es un Puntero?

Normalmente, una variable guarda un **valor** (como el número `25`). Un puntero, en cambio, guarda la **dirección de memoria** donde está ese valor.

- **`&` (Operador de dirección):** Se usa para obtener la dirección de una variable. ("¿Dónde vive?")
    
- **`*` (Operador de desreferencia):** Se usa para ver o cambiar el valor que hay en esa dirección. ("¿Qué hay dentro?")
    


```go
nombre := "Miguel"
puntero := &nombre // 'puntero' guarda la dirección de memoria de 'nombre'

fmt.Println(puntero)  // Imprime algo como: 0xc000010250
fmt.Println(*puntero) // Imprime: "Miguel" (el valor al que apunta)
```

---

### 2. Paso por Valor vs. Paso por Referencia

Aquí es donde está el truco para el **Alto Rendimiento**:

#### A. Paso por Valor (Copia)

Cuando pasas una variable a una función, Go crea una **copia** nueva. Si la función cambia el valor, la variable original **no cambia**.

- **Se usa en:** `int`, `float`, `string`, `bool` y **Arrays**.
    


```go
func cambiarValor(n int) {
    n = 100 // Esto solo cambia la copia
}
```

#### B. Paso por Referencia (Puntero)

Si pasas el puntero (`*int`), la función recibe la dirección real. Si la función cambia el valor, **la variable original sí cambia**.

- **Ventaja de rendimiento:** Si tienes un objeto gigante (como un archivo de 1GB), no lo copias; solo pasas la dirección (unos pocos bytes).
    


```go
func cambiarReal(n *int) {
    *n = 100 // Esto cambia el valor original en memoria
}
```

| **Tipo**                        | **Comportamiento**            | **Rendimiento**                     |
| ------------------------------- | ----------------------------- | ----------------------------------- |
| **Valor** (`int`, `Array`)      | Copia todo el dato.           | Lento si el dato es grande.         |
| **Referencia** (`Punteros`)     | Copia solo la dirección.      | **Rápido** (siempre pesa lo mismo). |
| **Híbridos** (`Slices`, `Maps`) | Se comportan como referencia. | Muy eficiente.                      |

### 🧠 El "Mindset" de Go:

A diferencia de otros lenguajes, en Go **usamos punteros principalmente por dos razones**:

1. Para que una función pueda **modificar** un dato original.
    
2. Por **eficiencia**, para no copiar estructuras de datos muy grandes en memoria.

# Resumen

**Punteros en 1 frase:** Es una variable que guarda la dirección de memoria de OTRA variable.

- `&variable` -> "Dime la dirección de esta variable" (Dirección).
    
- `*puntero` -> "Dime qué hay dentro de esta dirección" (Valor).

**¿Por qué en Backend?** Para modificar datos originales dentro de funciones sin crear copias pesadas en la RAM.
