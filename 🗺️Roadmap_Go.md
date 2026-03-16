
```dataviewjs
// 1. Buscamos todas las tareas en TODO el vault (puedes filtrar por carpeta si prefieres)
const todasLasTareas = dv.pages().file.tasks;

// 2. Filtramos para contar solo las que realmente son objetivos (puedes ajustar el filtro)
let total = todasLasTareas.length;
let completadas = todasLasTareas.where(t => t.completed).length;

// 3. Calculamos el porcentaje
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

// 4. Renderizado estético
dv.header(2, "🚀 Progreso Total: " + porcentaje + "%");
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 25px; accent-color: #00ADD8;"}});
dv.paragraph(`*Has completado **${completadas}** de **${total}** objetivos en tu camino al Alto Rendimiento.*`);
```

---

## 🌐 Módulo A: Fundamentos de Redes

```dataviewjs

const p = dv.current();
// Filtramos tareas que estén bajo el encabezado "Módulo A"
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Módulo A"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso Módulo A:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #29b6f6;"}});

```

- [ ] **Modelo OSI y TCP/IP:** Entender Capas 4 y 7.
    
- [ ] **Direccionamiento:** IP Estática/Dinámica, IPv4 vs IPv6, DHCP.
    
- [ ] **DNS:** Traducción de dominios a IP.
    
- [ ] **Puertos y Sockets:** La base de la comunicación.
    

---

## 🐹 Módulo B: Go para el Backend

```dataviewjs
// 1. Obtenemos el archivo actual y todas sus tareas
const currentFile = dv.current().file;
const allTasks = currentFile.tasks;

// 2. Definimos los límites del Módulo B mediante los encabezados
const cache = app.metadataCache.getFileCache(currentFile);
const startHeading = cache.headings.find(h => h.heading.includes("Módulo B"));
const endHeading = cache.headings.find(h => h.heading.includes("Módulo C"));

// 3. Obtenemos las líneas exactas de inicio y fin
const startLine = startHeading ? startHeading.position.start.line : 0;
const endLine = endHeading ? endHeading.position.start.line : 9999;

// 4. Filtramos las tareas que están físicamente entre esos dos encabezados
const tareasModuloB = allTasks.filter(t => t.line > startLine && t.line < endLine);

// 5. Cálculos
const total = tareasModuloB.length;
const completadas = tareasModuloB.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

// 6. Renderizado
dv.header(2, "📊 Progreso Detallado Módulo B: " + porcentaje + "%");
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 20px; accent-color: #00ADD8;"}});
dv.paragraph(`*Llevas **${completadas}** de **${total}** tareas completadas en el Módulo B.*`);
```

### 1. [[00_Indice Fundamentos|Fundamentos de Go]] 

```dataviewjs

const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Fundamentos de Go"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.1:** ${porcentaje}%`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00ADD8;"}});

```

- [x] [[01_Sintaxis, Tipado y Operadores]]-
    
- [x] [[02_Custom Types]]
    
- [x] [[03_Generics]]
    
- [x] [[04_Salida Formateada]]
    
- [x] [[05_ Entrada de datos (inputs)]]
    
- [x] [[06_Estructuras de Control]]
    
- [x] [[07_Punteros y Memoria (Valor vs Referencia)]]
    
- [x] [[08_Arrays, Slices y Maps]]
    

### 2. [[00_Indice Estructuras y Composición|Estructura y Composición]]

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Estructura y Composición"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.2:** ${porcentaje}%`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00ADD8;"}});
```

- [x] [[01_Funciones y Retornos Múltiples]]
    
- [x] [[02_Manejo de Errores y Control de Flujo]]
    
- [x] [[03_Structs y Methods]]
    
- [x] [[04_Interfaces]]
    
- [x] [[05_Embedding vs Herencia]]
    

### 3. Concurrencia Nativa

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Concurrencia Nativa"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.3:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00819c;"}});
```

- [ ] **Goroutines:** Procesos ligeros.
    
- [ ] **Channels y Select:** Comunicación segura.
    
- [ ] **Paquete sync:** `WaitGroup`, `Mutex`, `Once`.
    
- [ ] **Context:** Timeouts y cancelaciones.
    

### 4. Git y Flujo de Trabajo

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Git y Flujo de Trabajo"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.3:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00819c;"}});
```

- [ ] Control de Versiones y Pull Requests.
    
- [ ] **Conventional Commits:** Automatización de versiones.
    

### 5. Arquitectura y Diseño

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Arquitectura y Diseño"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.3:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00819c;"}});
```

- [ ] **Standard Library:** Dominar `net/http`.
    
- [ ] **Frameworks:** Echo o Fiber.
    
- [ ] **Arquitectura Hexagonal:** Separación de lógica de negocio.
    

### 6. Bases de Datos y Persistencia

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Bases de Datos y Persistencia"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.3:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00819c;"}});
```

- [ ] **PostgreSQL:** Diseño, índices y transacciones.
    
- [ ] **Drivers:** `pgx` o `sqlx` (Alto rendimiento).
    
- [ ] **Migraciones:** `golang-migrate`.
    

### 7. Comunicación Avanzada

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Comunicación Avanzada"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.7:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00819c;"}});
```

- [ ] **Protobuf:** Mensajes binarios.
    
- [ ] **gRPC:** Comunicación interna de alto rendimiento.
    

### 8. Testing y Calidad

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Testing y Calidad"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.8:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00819c;"}});
```

- [ ] Unit Testing y Table-Driven Tests.
    
- [ ] Mocking y Benchmarking.
    

### 9. Observabilidad y Profiling

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Observabilidad y Profiling"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.9:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00819c;"}});
```

- [ ] Logging estructurado (Zap) y Tracing.
    
- [ ] **pprof:** Identificación de cuellos de botella.
    

### 10. DevOps y Despliegue

```dataviewjs
const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("DevOps y Despliegue"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso B.10:** ${porcentaje}% (${completadas}/${total})`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 10px; accent-color: #00819c;"}});
```

- [ ] **Docker:** Imágenes multi-stage (binarios ligeros).
    
- [ ] CI/CD con GitHub Actions.
    

---

## 🚀 Módulo C: El Servidor y la "Vida Real"

```dataviewjs

const p = dv.current();
const tareas = p.file.tasks.filter(t => t.section.subpath && t.section.subpath.includes("Módulo C: El Servidor"));
const total = tareas.length;
const completadas = tareas.where(t => t.completed).length;
const porcentaje = total > 0 ? Math.round((completadas / total) * 100) : 0;

dv.paragraph(`**Progreso Módulo C:** ${porcentaje}%`);
dv.el("progress", "", {attr: {value: completadas, max: total, style: "width: 100%; height: 20px; accent-color: #ff5722;"}});

```

- [ ] **Administración Remota:** SSH (Llaves públicas/privadas).
    
- [ ] **Reverse Proxies:** Nginx o Caddy.
    
- [ ] **Seguridad:** Firewalls (UFW).
    
- [ ] **Self-Hosting:** Despliegue de DBs y binarios propios.
    
- [ ] **Kubernetes:** Escalabilidad básica.