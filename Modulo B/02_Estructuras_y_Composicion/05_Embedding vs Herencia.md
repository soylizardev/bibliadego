# Embedding vs Herencia

En Go, aunque puedes hacer cosas que _parecen_ POO, el paradigma se define oficialmente como **Programación Orientada a la Composición** (o simplemente un lenguaje **basado en tipos y composición**).

En Go, el concepto de **Herencia** (propio de Java/C++) no existe. En su lugar, utilizamos **Embedding** (Incrustación), que es la implementación práctica del principio: _"Prefiere la composición sobre la herencia"_.

---

## 1. La Filosofía

- **Java (Herencia):** Relación **"Es un"** (_is-a_). Un `Gato` **es un** `Animal`.
    
- **Go (Composición):** Relación **"Tiene un"** (_has-a_). Un `Gato` **tiene** características de `Animal`.
    

---

## 2. Cómo funciona el Embedding

Para "heredar" campos y métodos de otro struct, simplemente lo ponemos dentro de nuestro nuevo struct **sin darle un nombre de campo**.

```go
type Persona struct {
    Nombre string
    Edad   int
}

func (p Persona) Saludar() {
    fmt.Printf("Hola, soy %s\n", p.Nombre)
}

type Empleado struct {
    Persona // <--- Esto es Embedding (Campo anónimo)
    Sueldo  float64
}

func main() {
    e := Empleado{
        Persona: Persona{Nombre: "Miguel", Edad: 25},
        Sueldo:  1500.0,
    }

    // ¡Magia! Empleado puede usar los campos y métodos de Persona directamente
    fmt.Println(e.Nombre) 
    e.Saludar() 
}
```

---

## 3. Promoción de Métodos

Cuando haces embedding, los métodos del struct interno se **promueven** al struct externo. Es decir, `Empleado` gana automáticamente el método `Saludar()`.

---

## 4. Sobrescritura (Shadowing)

Si el struct externo define un método con el mismo nombre que el interno, Go usará el del externo. Esto es lo más parecido al `@Override` de Java.

```go
func (e Empleado) Saludar() {
    fmt.Printf("Hola, soy el empleado %s y gano %.2f\n", e.Nombre, e.Sueldo)
}
```

---

## 5. El problema de la Ambigüedad

Go permite "herencia múltiple" de facto mediante composición, pero si dos structs embebidos tienen el mismo método, Go se detiene y pide que tú definas cuál usar.

```go
type AllRounder struct {
    Striker // Tiene Pelear()
    Grappler // Tiene Pelear()
}

// SOLUCIÓN: Definir Pelear en AllRounder para quitar la ambigüedad
func (a AllRounder) Pelear() {
    a.Striker.Pelear() // Decides usar el de Striker
}
```

---

## 💡 Cuadro Comparativo Herencia vs Embedding

|**Característica**|**Herencia (Java)**|**Composición/Embedding (Go)**|
|---|---|---|
|**Jerarquía**|Vertical (Padres/Hijos)|Horizontal (Módulos/Piezas)|
|**Acoplamiento**|Fuerte (Frágil)|Débil (Flexible)|
|**Múltiple**|No permitida (Normalmente)|Permitida y común|
|**En tiempo de...**|Compilación (Estática)|Compilación (Pero más dinámica)|

---

📝 Términos que debes conocer:

- **Composición sobre Herencia (Composition over Inheritance):** Es el mantra de Go. En lugar de una jerarquía vertical (padre-hijo), usamos una estructura horizontal donde un struct contiene a otros para ganar sus habilidades.
    
- **Duck Typing (Tipado de Pato):** Es el nombre del sistema de interfaces de Go. _"Si camina como pato y grazna como pato, lo trato como un pato"_. No importa el origen del struct, solo que tenga el método necesario.
    
- **Polimorfismo Ad-hoc:** El polimorfismo en Go es mucho más dinámico porque puedes hacer que un struct sea compatible con una interfaz en cualquier momento, solo añadiendo el método, sin tocar la definición original del struct.
    

---

### 💡 Un dato curioso

El creador de Go, Rob Pike, dice que Go es un lenguaje **Post-POO**. Tomaron lo bueno de los objetos (encapsulamiento y métodos) y quitaron lo malo (la herencia compleja que a veces hace que el código sea un desastre de mantener).
