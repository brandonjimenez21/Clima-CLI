# 🌦️ Clima Desktop App v2.0 (Wails Edition)

[![Go Version](https://img.shields.io/badge/Go-1.24.2-blue?logo=go)](https://golang.org/)
[![Wails](https://img.shields.io/badge/Wails-v2-red)](https://wails.io/)
[![License](https://img.shields.io/badge/license-MIT-green)](https://opensource.org/licenses/MIT)

**Clima Desktop App** es la evolución de Clima CLI hacia una aplicación de escritorio nativa, moderna y profesional. Utiliza **Go** para una lógica de backend potente y eficiente, combinada con una interfaz de usuario web elegante y fluida gracias a **Wails**.

---

## ✨ Características Principales

- **Visuales Enriquecidos:** Temas dinámicos (Lluvia, Nubes, Despejado), ciclo día/noche y **Golden Hour** animado.
- **Alta Precisión:** Consultas optimizadas mediante coordenadas geográficas (Lat/Lon) para pronósticos de 24h exactos.
- **Inteligencia de Ubicación:** Auto-detección de ciudad mediante IP al arrancar la aplicación.
- **Personalización:** Sistema de favoritos persistente y métricas avanzadas (Lluvia, Amanecer/Atardecer).
- **Diseño Responsive:** Interfaz adaptativa 100% fluida para cualquier tamaño de ventana.

---

## 🛠️ Stack Tecnológico

- **Backend:** Go 1.24.2 (Lógica de API y procesamiento geográfico).
- **Frontend:** HTML5, CSS3 (Responsive), JavaScript (Vanilla).
- **Framework:** [Wails v2](https://wails.io/) (Binding nativo entre Go y JS).
- **APIs:** OpenWeatherMap, IPAPI (Geolocalización).

---

## 🏗️ Arquitectura de la Aplicación

La aplicación utiliza un puente de comunicación asíncrono y persistencia local:
1.  **Auto-Init:** El backend detecta la ubicación por IP al inicio para ofrecer datos inmediatos.
2.  **Precisión Geográfica:** Las búsquedas obtienen coordenadas exactas, usadas para pedir el pronóstico de 24h (Gráfica), garantizando la máxima precisión técnica.
3.  **Persistencia:** Los favoritos se guardan localmente para acceso rápido y sin latencia.
4.  **UI Adaptativa:** Estilos CSS fluidos con efectos visuales acelerados por hardware.

---

## 🚀 Instalación y Desarrollo

### Requisitos Previos
- **Go 1.24.2** o superior.
- **Wails CLI** (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`).
- **WebView2** (Requerido para Windows).

### Guía de Desarrollo
1.  **Configura tu API Key:** Asegúrate de tener tu archivo `.env` con `CLIMA_API_KEY`.
2.  **Modo Desarrollo:**
    ```bash
    wails dev
    ```
3.  **Compilar para Producción:**
    ```bash
    wails build
    ```

---

## 📂 Estructura del Proyecto

```text
clima-cli/
├── frontend/             # Interfaz Responsive y lógica JS
├── internal/             # Lógica de negocio (API con precisión Lat/Lon, Models)
├── main.go               # Punto de entrada de Wails
├── app.go                # Controlador (Backend Bridge con Geolocation)
└── wails.json            # Configuración del proyecto
```

---

## 🐳 Herencia CLI
Este proyecto nació como una herramienta de terminal (`cmd/clima-cli`) y conserva su lógica central robusta, ahora potenciada con una experiencia visual de escritorio de primer nivel.

---
Desarrollado con ❤️ combinando precisión de datos y diseño minimalista.
