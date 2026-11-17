# Implementación de una API REST para gestión de tareas

El sistema debe permitir la gestión de tareas a través de una API REST construida en Go utilizando el framework Gin. Se requiere un middleware propio de logging y un repositorio en memoria. El título de la tarea no puede estar vacío y debe tener un máximo de 100 caracteres.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | API REST de tareas con Go y Gin |
| **Nivel** | senior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Stack

- **Lenguaje:** Go
- **Framework:** Gin
- **Version:** 1.10.0

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Tener el proyecto base del reto funcionando en tu máquina: clonar el repositorio, instalar las herramientas y verificar que compila y arranca antes de empezar a practicar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Instala las herramientas para ejecutar el proyecto: Go 1.21+.
- Clona el repositorio del reto (el enlace viene en la entrega del reto y en la sección "Repositorio del Proyecto" del PDF) y ábrelo en tu IDE.
- Instala las dependencias y compila ejecutando `go mod tidy && go build ./...` en la raíz del repositorio.
- Confirma que el comando termina sin errores: el proyecto arranca y sus pruebas corren.
- Ubica los archivos de práctica: son los que contienen `TODO(reto)` (búscalos en el repositorio). Lee el README.md y el AGENTS.md antes de empezar.

**Entregable:** El proyecto compila/arranca sin errores con `go mod tidy && go build ./...`.

<details>
<summary>Pistas de conocimiento</summary>

- Los archivos con `TODO(reto)` son el ejercicio: el resto del repositorio ya está resuelto y no hace falta tocarlo.
- Si el comando de verificación falla, revisa la versión de las herramientas instaladas antes de modificar código.
- AGENTS.md describe el reto, el comando de verificación y las reglas del repositorio.

</details>

### Fase 1: Configuración inicial y endpoint de creación de tareas

**Objetivo:** Configurar el proyecto y crear el endpoint para la creación de tareas.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Configurar el proyecto en Go utilizando Gin.
- Implementar un middleware de logging personalizado.
- Crear un endpoint POST para la creación de tareas, validando que el título no esté vacío y tenga un máximo de 100 caracteres.

**Entregable:** Proyecto de Go con Gin configurado, middleware de logging implementado y endpoint POST para creación de tareas.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo estructurar tu proyecto para una fácil mantenibilidad.
- Piensa en cómo manejar los errores de validación de forma efectiva.

</details>

### Fase 2: Endpoint para listar tareas

**Objetivo:** Implementar un endpoint GET para listar todas las tareas creadas.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Crear un endpoint GET que devuelva una lista de todas las tareas almacenadas en el repositorio en memoria.

**Entregable:** Endpoint GET implementado que devuelve una lista de todas las tareas.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo estructurar la respuesta del endpoint para una fácil comprensión y uso.
- Piensa en cómo manejar la ausencia de tareas en el repositorio.

</details>

### Fase 3: Endpoint para actualizar tareas

**Objetivo:** Implementar un endpoint PUT para actualizar una tarea existente.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Crear un endpoint PUT que permita actualizar el título de una tarea existente, validando que el nuevo título cumpla con las restricciones de validación.

**Entregable:** Endpoint PUT implementado para actualizar el título de una tarea.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo identificar de manera única una tarea para su actualización.
- Piensa en cómo manejar los errores de actualización, como la tarea no encontrada.

</details>

### Fase 4: Revisión y optimización

**Objetivo:** Revisar y optimizar el código implementado en las fases anteriores.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Revisar el código implementado para identificar posibles mejoras y optimizaciones.
- Aplicar las mejoras identificadas al código.

**Entregable:** Código optimizado y mejorado de la API REST.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo mejorar la legibilidad y mantenibilidad del código.
- Piensa en posibles optimizaciones de rendimiento.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es una API REST y cuál es su propósito en este reto?
- **paraQueSirve**: ¿Para qué sirve el middleware de logging en este reto?
- **comoSeUsa**: ¿Cómo se usa el endpoint PUT para actualizar una tarea en este reto?
- **erroresComunes**: ¿Cuáles son los errores comunes que pueden ocurrir al validar el título de una tarea en este reto?
- **queDecisionesImplica**: ¿Qué decisiones implica la revisión y optimización del código en este reto?

## Criterios de Evaluacion

- Configuración inicial y endpoint de creación de tareas implementados correctamente.
- Endpoint GET para listar tareas implementado correctamente.
- Endpoint PUT para actualizar tareas implementado correctamente.
- Middleware de logging personalizado implementado correctamente.
- Validación del título de la tarea según las restricciones implementada correctamente.
- Código revisado y optimizado correctamente.

## Como trabajar con un asistente de IA

Hay dos caminos, elegi uno:

- **AGENTS.md** (recomendado) — instrucciones nativas del repo. Abri esta carpeta con tu agente local (Claude Code, Cursor, Codex, Copilot, Gemini) y las carga solo. Sabe que archivos faltan y con que comando se verifica, y completa el scaffold escribiendo en disco.
- **PROMPT_MEJORA.md** — para copiar y pegar en un chat (claude.ai, ChatGPT) si no tenes un agente en el IDE. Declara lo mismo que AGENTS.md; los archivos que te devuelva el chat los guardas en esta carpeta respetando las rutas.

Ninguno de los dos resuelve las fases del reto: eso es tu trabajo.

## Verificacion

El proyecto esta listo para trabajar cuando este comando corre sin errores:

```bash
go mod tidy && go build ./...
```

---

*Reto generado automaticamente por Challenge Generator - Pragma*
