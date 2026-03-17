# Conversión de Datos (Paquete strconv)

Cuando capturamos datos con `bufio`, Go los recibe como un `string`. Para realizar operaciones matemáticas, debemos convertirlos a tipos numéricos (`int`, `float`, etc.) utilizando el paquete **`strconv`**.

## 1. Convertir String a Entero (`ParseInt`)

Es la función más completa para convertir texto en números enteros.

```go
numero, err := strconv.ParseInt(texto, base, bitSize)
```

### Parámetros técnicos:

- **`s` (string):** El texto que quieres convertir (ej: `"100"`).
    
- **`base` (int):** El sistema numérico del texto.
    
    - `10`: Decimal (el más usado).
        
    - `2`: Binario.
        
    - `16`: Hexadecimal.
        
    - `0`: Autodetectar (según prefijos como `0x`).
        
- **`bitSize` (int):** El tamaño del tipo de destino.
    
    - `0, 8, 16, 32, 64`. Define si el número debe caber en un `int8`, `int16`, etc.
        
    - **Nota:** Aunque definas un tamaño menor, la función siempre devuelve un tipo `int64`.
        

---

## 2. La vía rápida: `Atoi` (S-to-i)

Si solo necesitas convertir un string a un entero común (`int`) en base 10, existe una función simplificada llamada `Atoi` (_String to Integer_).

```go
texto := "25"
numero, err := strconv.Atoi(texto)

// Es el equivalente exacto a:
// strconv.ParseInt(texto, 10, 0)
```

---

## 3. Convertir String a Decimal (`ParseFloat`)

Para números con punto decimal, usamos `ParseFloat`.

```go
texto := "3.1416"
// El segundo parámetro es el bitSize (32 o 64)
decimal, err := strconv.ParseFloat(texto, 64)
```

---

## 4. El camino inverso: Número a String

A veces necesitas convertir un número en texto (por ejemplo, para concatenarlo en un mensaje o guardarlo en un archivo).

### De Entero a String (`Itoa`)

```go
edad := 30
texto := strconv.Itoa(edad) // "Integer to ASCII"
```

### De Decimal a String (`FormatFloat`)

Es más complejo porque requiere definir el formato y la precisión.

```go
valor := 12.3456
// 'f' significa sin exponente, 2 es la cantidad de decimales, 64 es el tamaño
texto := strconv.FormatFloat(valor, 'f', 2, 64) // Resultado: "12.35"
```

---

## 💡 Tips:

> [!IMPORTANT] **Manejo de Tipos:** Recuerda que `ParseInt` siempre devuelve un `int64`. Si tu función espera un `int` normal, deberás convertirlo manualmente después de la función: `resultadoFinal := int(numeroConvertido)`

> [!WARNING] **Limpieza previa:** Antes de usar `strconv`, asegúrate de usar `strings.TrimSpace(texto)` para eliminar saltos de línea (`\n`) o espacios accidentales, de lo contrario la conversión fallará con un error.

---