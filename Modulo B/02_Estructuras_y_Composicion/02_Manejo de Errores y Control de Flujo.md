En Go, el manejo de errores es **explícito**. No existen las excepciones (`try-catch`) como en Java. Esto nos obliga a tratar los errores como parte de la lógica de negocio, lo que hace que el software de backend sea extremadamente robusto y predecible.

## 1. El Patrón Estándar: `if err != nil`

Las funciones que pueden fallar siempre devuelven un tipo `error` como último valor de retorno.

```go
func buscarUsuario(id int) (User, error) {
    // ... lógica de búsqueda
    if noEncontrado {
        return User{}, fmt.Errorf("usuario con id %d no existe", id)
    }
    return usuario, nil // nil significa que no hubo error
}

// Uso profesional:
user, err := buscarUsuario(10)
if err != nil {
    // Aquí decides qué hacer: log, retornar el error o morir
    fmt.Println("Error encontrado:", err)
    return 
}
fmt.Println("Usuario encontrado:", user.Nombre)
```

## 2. `defer`: El "Limpiador" Automático

La palabra clave `defer` asegura que una función se ejecute justo antes de salir de la función actual. Es vital para la limpieza de recursos y evitar fugas de memoria (_memory leaks_).

- **Uso común:** Cerrar archivos, cerrar conexiones a DB, liberar bloqueos (Mutex).
    
- **Orden de ejecución:** Si hay varios `defer`, se ejecutan en orden **LIFO** (Last In, First Out), como una pila de platos.

```go
f, err := os.Open("config.json")
if err != nil {
    return err
}
defer f.Close() // Se ejecutará al final de la función, pase lo que pase
```

## 3. Errores Críticos: `panic` y `recover`

- **`panic`**: Detiene la ejecución normal del programa inmediatamente. Solo se debe usar en errores irrecuperables (ej: no se pudo cargar la variable de entorno de la DB al arrancar).
    
- **`recover`**: Solo funciona dentro de una función llamada por `defer`. Permite capturar un `panic` para que el programa no "explote" y pueda seguir vivo.
    

```go
func rescatar() {
    if r := recover(); r != nil {
        fmt.Println("Recuperado de un error crítico:", r)
    }
}

func miFuncion() {
    defer rescatar()
    panic("¡Conexión perdida con el satélite!")
}
```

## 4. Mejores Prácticas para tu Roadmap

1. **No ignores los errores:** Nunca uses `_ , _ = funcion()`. Siempre verifica el `err`.
    
2. **Los errores son valores:** Puedes compararlos y envolverlos (`fmt.Errorf("... %w", err)`) para dar más contexto.
    
3. **Defer temprano:** Pon el `defer` inmediatamente después de comprobar que el recurso se abrió correctamente.