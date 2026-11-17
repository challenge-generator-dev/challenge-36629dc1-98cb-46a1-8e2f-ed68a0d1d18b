# Prompt para Mejorar el Codigo Base

Copia y pega el contenido del bloque de abajo en un asistente de IA (Claude, ChatGPT)
para obtener el proyecto completo y arrancable. Guarda los archivos que te devuelva en la
carpeta del repo, respetando exactamente las rutas.

Si preferis trabajar en tu editor con un agente local (Claude Code, Cursor, Copilot), usa `AGENTS.md` en vez de este archivo: dice lo mismo pero para que escriba los archivos en disco.

## Las dos reglas que no se negocian

1. **Completa el boilerplate.** Todo lo que el proyecto necesita para compilar y arrancar: manifiesto de dependencias, punto de entrada, configuracion, capa de interfaz, y las capas del patron arquitectonico declarado. Eso es andamiaje y es tu trabajo.
2. **NO resuelvas el reto.** Los entregables de las fases son el trabajo de la persona. El hueco pedagogico se deja como esta: el proyecto arranca, pero lo que el reto pide implementar NO esta implementado.

Dicho de otra forma: si algo impide compilar, arreglalo. Si algo es logica de negocio incompleta, validaciones ausentes, un secreto hardcodeado o un patron mejorable, dejalo exactamente como esta — es lo que la persona tiene que encontrar.

## Superficie de practica — NO resuelvas

Estos archivos SON el ejercicio de la persona. No los implementes; deja stubs.

- `internal/domain/infrastructure/middleware/logging_middleware.go` — La fase "Configuración inicial y endpoint de creación de tareas" ya venia resuelta por el codigo base: este archivo pasa a ser el ejercicio.

## Como saber que terminaste

```bash
go mod tidy && go build ./...
```

Ese comando corriendo sin errores es la definicion de "listo".

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Perfil
Chapter Backend, Especialidad Desarrollador, Tecnología Go, Senior L2

### Brecha de conocimiento
Construir una API REST en Go con Gin para gestionar tareas

### Misión / candidato
Exponer la gestion de tareas por HTTP

### Reto
- Tema: API REST de tareas con Go y Gin
- Seniority: senior-l2
- Tipo: practical
- Título: Implementación de una API REST para gestión de tareas
- Tiempo estimado: 8 horas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Configuración inicial y endpoint de creación de tareas — objetivo: Configurar el proyecto y crear el endpoint para la creación de tareas. — entregable (NO resolver): Proyecto de Go con Gin configurado, middleware de logging implementado y endpoint POST para creación de tareas.
- Fase 2: Endpoint para listar tareas — objetivo: Implementar un endpoint GET para listar todas las tareas creadas. — entregable (NO resolver): Endpoint GET implementado que devuelve una lista de todas las tareas.
- Fase 3: Endpoint para actualizar tareas — objetivo: Implementar un endpoint PUT para actualizar una tarea existente. — entregable (NO resolver): Endpoint PUT implementado para actualizar el título de una tarea.
- Fase 4: Revisión y optimización — objetivo: Revisar y optimizar el código implementado en las fases anteriores. — entregable (NO resolver): Código optimizado y mejorado de la API REST.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (entrega de archivos).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Entrega de archivos
Entregá cada archivo con su ruta completa, respetando exactamente la estructura de rutas
indicada por los marcadores: un bloque por archivo, precedido por su ruta. No generes ZIPs
ni archivos adjuntos.
La entrega debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: Entregá los archivos completos al finalizar. No preguntes si el usuario
quiere recibirlos. Simplemente entregalos; No debes desplegar en el chat el resumen de lo que arreglaste, solo entregá los archivos.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del último archivo entregado
No preguntes si el usuario quiere los archivos: simplemente entregalos siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: go.mod ===
module github.com/user/project

go 1.21

require (
	github.com/gin-gonic/gin v1.9.1
)

// === ARCHIVO: main.go ===
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/middleware"
	"github.com/user/project/internal/domain/infrastructure/repository"
	"github.com/user/project/internal/domain/infrastructure/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := application.NewTaskService(taskRepo)

	r := router.SetupRouter()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}

// === ARCHIVO: internal/app/routes.go ===
package app

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	// Sample route - will be replaced by domain code
	router.GET("/sample", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sample endpoint",
		})
	})
}

// === ARCHIVO: internal/contracts/types.go ===
package contracts

// Sample contract types - interfaces and models go here

// === ARCHIVO: internal/domain/cmd/api/main.go ===
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/handler"
	"github.com/user/project/internal/domain/infrastructure/middleware"
	"github.com/user/project/internal/domain/infrastructure/repository"
	"github.com/user/project/internal/domain/infrastructure/router"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	taskRepo := repository.NewInMemoryTaskRepository()
	taskService := application.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	engine := router.SetupRouter()

	engine.Use(gin.Recovery())
	engine.Use(middleware.LoggingMiddleware())

	engine.POST("/tasks", func(c *gin.Context) {
		taskHandler.CreateTask(c)
	})

	engine.GET("/tasks", func(c *gin.Context) {
		taskHandler.ListTasks(c)
	})

	engine.PUT("/tasks/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
			return
		}
		c.Set("taskID", id)
		taskHandler.UpdateTask(c)
	})

	engine.PATCH("/tasks/:id/complete", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
			return
		}
		c.Set("taskID", id)
		taskHandler.CompleteTask(c)
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Println("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %v\n", err)
	}
}

// === ARCHIVO: internal/domain/task.go ===
// TODO(reto): Implementar la entidad Task con sus reglas de dominio.
// Requisitos:
// - El título no puede estar vacío y debe tener máximo 100 caracteres.
// - El ID debe ser generado automáticamente como UUID.
// - La tarea debe tener un estado (pendiente, completada).
package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEmptyTitle    = errors.New("el título no puede estar vacío")
	ErrTitleTooLong  = errors.New("el título no puede exceder 100 caracteres")
	ErrInvalidStatus = errors.New("estado inválido")
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "pendiente"
	StatusCompleted TaskStatus = "completada"
)

type Task struct {
	ID        uuid.UUID   `json:"id"`
	Title     string      `json:"title"`
	Status    TaskStatus  `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// TODO(reto): Implementar NewTask para crear una nueva tarea con validaciones.
func NewTask(title string) (*Task, error) {
	return nil, nil
}

// TODO(reto): Implementar UpdateTitle para actualizar el título con validaciones.
func (t *Task) UpdateTitle(newTitle string) error {
	return nil
}

// TODO(reto): Implementar Complete para marcar la tarea como completada.
func (t *Task) Complete() error {
	return nil
}

// === ARCHIVO: internal/domain/application/task_service.go ===
package application

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/user/project/internal/domain"
	"github.com/user/project/internal/domain/infrastructure/repository"
)

var (
	ErrTaskNotFound  = errors.New("tarea no encontrada")
	ErrEmptyTitle    = errors.New("el título no puede estar vacío")
	ErrTitleTooLong  = errors.New("el título no puede exceder 100 caracteres")
)

const maxTitleLength = 100

type TaskService struct {
	repo *repository.InMemoryTaskRepository
}

func NewTaskService(repo *repository.InMemoryTaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) (*domain.Task, error) {
	trimmedTitle := strings.TrimSpace(title)
	if trimmedTitle == "" {
		return nil, ErrEmptyTitle
	}
	if len(trimmedTitle) > maxTitleLength {
		return nil, ErrTitleTooLong
	}

	task, err := domain.NewTask(trimmedTitle)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetAllTasks() ([]*domain.Task, error) {
	tasks, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) UpdateTaskTitle(id uuid.UUID, newTitle string) (*domain.Task, error) {
	trimmedTitle := strings.TrimSpace(newTitle)
	if trimmedTitle == "" {
		return nil, ErrEmptyTitle
	}
	if len(trimmedTitle) > maxTitleLength {
		return nil, ErrTitleTooLong
	}

	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if err := task.UpdateTitle(trimmedTitle); err != nil {
		return nil, err
	}

	if err := s.repo.Save(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) CompleteTask(id uuid.UUID) (*domain.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if err := task.Complete(); err != nil {
		return nil, err
	}

	if err := s.repo.Save(task); err != nil {
		return nil, err
	}

	return task, nil
}

// === ARCHIVO: internal/domain/infrastructure/repository/task_repository.go ===
// TODO(reto): Implementar el repositorio en memoria para tareas.
// Requisitos:
// - Almacenar tareas en un mapa con UUID como clave.
// - Implementar métodos para guardar, listar y buscar tareas.
// - Manejar concurrencia con sync.RWMutex.
package repository

import (
	"sync"

	"github.com/google/uuid"
	"github.com/user/project/internal/domain"
)

type InMemoryTaskRepository struct {
	tasks map[uuid.UUID]*domain.Task
	mu    sync.RWMutex
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[uuid.UUID]*domain.Task),
	}
}

// TODO(reto): Implementar Save para guardar una tarea.
func (r *InMemoryTaskRepository) Save(task *domain.Task) error {
	return nil
}

// TODO(reto): Implementar FindAll para listar todas las tareas.
func (r *InMemoryTaskRepository) FindAll() ([]*domain.Task, error) {
	return nil, nil
}

// TODO(reto): Implementar FindByID para buscar una tarea por ID.
func (r *InMemoryTaskRepository) FindByID(id uuid.UUID) (*domain.Task, error) {
	return nil, nil
}

// === ARCHIVO: internal/domain/infrastructure/middleware/logging_middleware.go ===
package middleware

import (
	"github.com/gin-gonic/gin"
)

// LoggingMiddleware registra las peticiones HTTP entrantes con detalles de método, ruta y tiempo de respuesta.
func LoggingMiddleware() gin.HandlerFunc {
	panic("TODO(reto): import")
}

// === ARCHIVO: internal/domain/dto/task_dto.go ===
package dto

import (
	"time"

	"github.com/google/uuid"
)

type TaskCreateDTO struct {
	Title string `json:"title" binding:"required,max=100"`
}

type TaskUpdateDTO struct {
	Title string `json:"title" binding:"required,max=100"`
}

type TaskResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTaskResponseDTO(id uuid.UUID, title string, status string, createdAt, updatedAt time.Time) TaskResponseDTO {
	return TaskResponseDTO{
		ID:        id,
		Title:     title,
		Status:    status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

// === ARCHIVO: internal/domain/infrastructure/handler/task_handler.go ===
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
)

type TaskHandler struct {
	service *application.TaskService
}

func NewTaskHandler(service *application.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// CreateTask handles POST /tasks
// TODO(reto): Implementar la lógica de creación de tareas
func (h *TaskHandler) CreateTask(c *gin.Context) {
	panic("TODO(reto): CreateTask")
}

// ListTasks handles GET /tasks
// TODO(reto): Implementar la lógica de listado de tareas
func (h *TaskHandler) ListTasks(c *gin.Context) {
	panic("TODO(reto): ListTasks")
}

// UpdateTask handles PUT /tasks/:id
// TODO(reto): Implementar la lógica de actualización de tareas
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	panic("TODO(reto): UpdateTask")
}

// CompleteTask handles PATCH /tasks/:id/complete
// TODO(reto): Implementar la lógica de completado de tareas
func (h *TaskHandler) CompleteTask(c *gin.Context) {
	panic("TODO(reto): CompleteTask")
}

// === ARCHIVO: internal/domain/infrastructure/router/router.go ===
// TODO(reto): archivo stubbeado; implementa los cuerpos.
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/user/project/internal/domain/application"
	"github.com/user/project/internal/domain/infrastructure/handler"
	"github.com/user/project/internal/domain/infrastructure/middleware"
	"github.com/user/project/internal/domain/infrastructure/repository"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Aplicar middleware de logging global
	router.Use(middleware.LoggingMiddleware())

	// Inicializar repositorio
	taskRepo := repository.NewInMemoryTaskRepository()

	// Inicializar servicio
	taskService := application.NewTaskService(taskRepo)

	// Inicializar handler
	taskHandler := handler.NewTaskHandler(taskService)

	// Definir rutas
	tasks := router.Group("/tasks")
	{
		tasks.POST("", taskHandler.CreateTask)
		tasks.GET("", taskHandler.ListTasks)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.PATCH("/:id/complete", taskHandler.CompleteTask)
	}

	return router
}
```
