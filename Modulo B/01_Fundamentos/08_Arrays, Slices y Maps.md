Take
Esta nota cubre los fundamentos de la gestión de colecciones de datos en Go, enfocada en el **Módulo B.1** de nuestra Biblia de Estudio.

# SECCIÓN 1: ARRAYS (Arreglos Estáticos)

**Definición:** Tienen un tamaño fijo que no puede cambiar. Se usan para listas donde sabes exactamente cuántos elementos habrá (ej. días de la semana).

## **Declaración y creación:**

``` go
- var miArray [3]int //Crea un array de 3 enteros, todos en 0
    
- numeros := [3]int{10, 20, 30} //Crea un array con valores iniciales
```

## **Operaciones:**

- **Leer:** `dato := numeros[0]`
    
- **Modificar:** `numeros[1] = 50`
    
- **Longitud:** `len(numeros)`

---

## SECCIÓN 2: SLICES (Arreglos Dinámicos)


**Definición:** Son punteros a un array. Pueden crecer y son mucho más eficientes para el Backend. Tienen "Pointer", "Len" (longitud) y "Cap" (capacidad).

### **Formas de crear un Slice:**

1. **Var (Nil):** `var lista []string` (No ocupa memoria todavía)
    
2. **Literal:** `lista := []string{"A", "B"}`
    
3. **Make (Alto Rendimiento):** `lista := make([]string, 0, 10)` (Crea un slice de tamaño 0 pero reserva espacio para 10 elementos de una vez).

### **Operaciones:**

- **Agregar (Create):** `lista = append(lista, "Nuevo Dato")`
    
- **Unir dos slices:** `lista = append(lista, otraLista...)`
    
- **Leer/Modificar:** Igual que el array: `lista[0] = "Editado"`
    
- **Borrar un elemento (índice i):** `lista = append(lista[:i], lista[i+1:]...)`

---

### SECCIÓN 3: MAPS (Diccionarios Clave-Valor)

**Definición:** Almacenan datos usando una clave única (como un ID o un nombre). Son ultra rápidos para buscar información.

**Formas de crear un Map:**

1. **Make:** `miMap := make(map[string]int)`
    
2. **Literal:** `edades := map[string]int{"Miguel": 25, "Jose": 30}`
    

**Operaciones:**

- **Agregar o Modificar:** `miMap["Clave"] = 100`
    
- **Borrar:** `delete(miMap, "Clave")`
    
- **Lectura Segura (Comma OK):**
    
	```go
	valor, ok := mapa[clave]
	if ok {
	    // La clave existe, usar 'valor'
	} else {
	    // La clave NO existe
	}
	```
cortar---
### SECCIÓN 4: NOTAS DE RENDIMIENTO (Para tu Roadmap)

1. **Pre-asignar:** Siempre que uses `make`, trata de ponerle una capacidad. Si no lo haces, Go tiene que crear nuevos arrays internos cada vez que el slice se llena, y eso consume CPU.
    
2. **Copia de datos:** Los Arrays se copian enteros (lento). Los Slices y Maps solo copian su "encabezado" (rápido). Por eso en Backend casi siempre usamos Slices.