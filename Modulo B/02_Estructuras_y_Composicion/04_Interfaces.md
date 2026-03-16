# Interfaces

Las interfaces nos permiten escribir código más flexible y re-utilizable sin depender de tipos en específicos. En Go, las **Interfaces** definen el **comportamiento** de un objeto. A diferencia de Java, no usas la palabra `implements`. Si un struct tiene los métodos que pide la interfaz, la implementa automáticamente (**Duck Typing**).

> _"Si camina como pato y grazna como pato, entonces es un pato"._

---

## 1. Definición de una Interfaz

Una interfaz solo contiene las **firmas** de los métodos (nombre, parámetros y retorno).

```go
type Mensajero interface {
    Enviar(mensaje string) error
}
```

---

## 2. Implementación Implícita

No necesitas decir nada en el struct. Solo crea el método con el **mismo nombre y firma**.

```go
type Email struct {
    Direccion string
}

// Email implementa Mensajero automáticamente por tener este método
func (e Email) Enviar(m string) error {
    fmt.Printf("Enviando Email a %s: %s\n", e.Direccion, m)
    return nil
}

type SMS struct {
    Numero string
}

// SMS también implementa Mensajero
func (s SMS) Enviar(m string) error {
    fmt.Printf("Enviando SMS al %s: %s\n", s.Numero, m)
    return nil
}
```

---

## 3. El Poder del Polimorfismo

Puedes crear funciones que reciban la **interfaz**, y así aceptarán cualquier struct que la cumpla. Esto permite **desacoplar** el código.

```go
func Notificar(m Mensajero, texto string) {
    m.Enviar(texto)
}

func main() {
    e := Email{Direccion: "miguel@go.dev"}
    s := SMS{Numero: "+58412000"}

    // La función acepta ambos porque ambos son 'Mensajeros'
    Notificar(e, "¡Hola desde Go!")
    Notificar(s, "¡Hola desde Go!")
}
```

---

## 4. La Interfaz Vacía (`interface{}` o `any`)

Una interfaz sin métodos es cumplida por **absolutamente todo**. Es el equivalente a `Object` en Java, pero en Go moderno (1.18+) usamos el alias `any`.

```go
func Describir(i any) {
    fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
}
```

_⚠️ Nota: No abuses de `any`, pierde la seguridad del tipado._

---

## 💡 Diferencias Clave con Java

|**Característica**|**Java**|**Go**|
|---|---|---|
|**Vínculo**|Explícito (`implements`)|**Implícito** (Automático)|
|**Contrato**|Rígido|Flexible (Desacoplado)|
|**Diseño**|Interfaces grandes|**Interfaces pequeñas** (1 o 2 métodos)|

---
