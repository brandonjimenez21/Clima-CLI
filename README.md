# 🌦️ Clima CLI v1.0

[![Go Version](https://img.shields.io/badge/Go-1.24-blue?logo=go)](https://golang.org/)
[![Docker Image](https://img.shields.io/badge/docker-clima--cli:2.0-blue?logo=docker)](https://hub.docker.com/)
[![License](https://img.shields.io/badge/license-MIT-green)](https://opensource.org/licenses/MIT)

**Clima CLI** es una aplicación de terminal (TUI) moderna, interactiva y minimalista escrita en **Go**. Convierte tu terminal en una estación meteorológica profesional con un solo clic.

---

## ✨ Características Principales

- **Interfaz Interactiva (TUI):** Construida con [Bubble Tea](https://github.com/charmbracelet/bubbletea), permite buscar múltiples ciudades sin reiniciar la app.
- **Sensación Térmica (Thermal Sensation):** Cálculo inteligente del impacto térmico real, categorizado humanamente (Gélido, Agradable, Caluroso, etc.).
- **Persistencia de Configuración:** Carga automática de la API Key desde archivos `.env`.
- **Modo "App":** Incluye lanzadores (`.bat` y `.sh`) para abrir la herramienta en su propia ventana dedicada.
- **Arquitectura Stateless:** Procesamiento 100% en memoria, sin bases de datos, garantizando rapidez y privacidad.

---

## 🛠️ Stack Tecnológico

- **Lenguaje:** Go 1.24 (aprovechando las últimas optimizaciones de rendimiento).
- **TUI Framework:** [Charmbracelet](https://charm.sh/) (Bubble Tea, LipGloss, Bubbles).
- **Contenerización:** Docker (Multi-stage build basado en Alpine).
- **CI/CD:** GitHub Actions (Lint, Test & Docker Build).

---

## 🚀 Instalación y Uso de un Solo Clic

### 1. Configuración Inicial
Obtén tu API Key gratuita en [OpenWeatherMap](https://openweathermap.org/api) y configúrala en el archivo `.env`:
```bash
# Crea tu archivo .env basado en el ejemplo
echo "CLIMA_API_KEY=tu_api_key_aqui" > .env
```

### 2. Construir la Imagen (Solo una vez)
```bash
docker build -t clima-cli:2.0 .
```

### 3. Ejecutar como una App
- **Windows:** Haz doble clic en `clima-app.bat`.
- **Linux/macOS:** Ejecuta `./clima-app.sh`.

---

## 💻 Desarrollo Local (Sin Docker)

Si prefieres ejecutarlo nativamente con Go:

1. **Instalar dependencias:**
   ```bash
   go mod tidy
   ```
2. **Ejecutar modo interactivo:**
   ```bash
   go run cmd/clima-cli/main.go
   ```
3. **Ejecutar comando directo:**
   ```bash
   go run cmd/clima-cli/main.go get "Madrid"
   ```

---

## 📂 Estructura del Proyecto

```text
clima-cli/
├── cmd/clima-cli/main.go     # Punto de entrada y orquestación
├── internal/
│   ├── api/client.go         # Consumo de API y lógica de entorno
│   ├── model/weather.go      # Estructuras y lógica de sensación térmica
│   └── ui/                   # Componentes visuales
│       ├── formatter.go      # Estilos LipGloss
│       └── tui.go            # Lógica interactiva Bubble Tea
├── .github/workflows/ci.yml  # Pipeline de automatización
├── clima-app.bat             # Lanzador para Windows
├── clima-app.sh              # Lanzador para Linux/macOS
├── Dockerfile                # Construcción optimizada Go 1.24
└── .env                      # Configuración privada (No subir al repo)
```

---

## 🐳 Detalles de Docker

La imagen está optimizada para ser ultra ligera (~20MB):
- **Stage 1 (Build):** Usa `golang:1.24-alpine` para compilar un binario estático.
- **Stage 2 (Final):** Usa `alpine:3.19` con certificados CA para seguridad en peticiones HTTPS.

---

## 📝 Notas de Versión (v1.0)
- Primera versión estable con soporte TUI completo.
- Validación de API Key y manejo de errores estilizado.
- Soporte para espacios en nombres de ciudades y caracteres especiales.

---
Desarrollado con ❤️ usando **Go** y **Charmbracelet**.
