# Biblia de estudio de Lizardev Backend/Go

---

## 🌐 Módulo A: Fundamentos de Redes

- **Modelo OSI y TCP/IP:** Entender Capas 4 y 7.
    
- **Direccionamiento:** IP Estática/Dinámica, IPv4 vs IPv6, DHCP.
    
- **DNS:** Traducción de dominios a IP.
    
- **Puertos y Sockets:** La base de la comunicación.
    

---

## 🐹 Módulo B: Go para el Backend

### 1. [[00_Indice Fundamentos|Fundamentos de Go]] 

- [[01_Sintaxis, Tipado y Operadores]]-
    
- [[02_Custom Types]]
    
- [[03_Generics]]
    
- [[04_Salida Formateada]]
    
- [[05_ Entrada de datos (inputs)]]
    
- [[06_Estructuras de Control]]
    
- [[07_Punteros y Memoria (Valor vs Referencia)]]
    
- [[08_Arrays, Slices y Maps]]
    

### 2. [[00_Indice Estructuras y Composición|Estructura y Composición]]

- [[01_Funciones y Retornos Múltiples]]
    
- [[02_Manejo de Errores y Control de Flujo]]
    
- [[03_Structs y Methods]]
    
- [[04_Interfaces]]
    
- [[05_Embedding vs Herencia]]
    

### 3. Concurrencia Nativa

- **Goroutines:** Procesos ligeros.
    
- **Channels y Select:** Comunicación segura.
    
- **Paquete sync:** `WaitGroup`, `Mutex`, `Once`.
    
- **Context:** Timeouts y cancelaciones.
    

### 4. Git y Flujo de Trabajo

- **Control de Versiones y Pull Requests.**
    
- **Conventional Commits:** Automatización de versiones.
    

### 5. Arquitectura y Diseño

- **Standard Library:** Dominar `net/http`.
    
- **Frameworks:** Echo o Fiber.
    
- **Arquitectura Hexagonal:** Separación de lógica de negocio.
    

### 6. Bases de Datos y Persistencia

- **PostgreSQL:** Diseño, índices y transacciones.
    
- **Drivers:** `pgx` o `sqlx` (Alto rendimiento).
    
- **Migraciones:** `golang-migrate`.
    

### 7. Comunicación Avanzada

- **Protobuf:** Mensajes binarios.
    
- **gRPC:** Comunicación interna de alto rendimiento.
    

### 8. Testing y Calidad

- **Unit Testing y Table-Driven Tests.**
    
- **Mocking y Benchmarking.**
    

### 9. Observabilidad y Profiling

- **Logging estructurado (Zap) y Tracing**.
    
- **pprof:** Identificación de cuellos de botella.
    

### 10. DevOps y Despliegue

- **Docker:** Imágenes multi-stage (binarios ligeros).
    
- **CI/CD con GitHub Actions.**
    

---

## 🚀 Módulo C: El Servidor y la "Vida Real"

- **Administración Remota:** SSH (Llaves públicas/privadas).
    
- **Reverse Proxies:** Nginx o Caddy.
    
- **Seguridad:** Firewalls (UFW).
    
- **Self-Hosting:** Despliegue de DBs y binarios propios.
    
- **Kubernetes:** Escalabilidad básica.
