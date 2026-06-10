@echo off
SETLOCAL
:: Obtener la ruta del directorio actual
SET "CURRENT_DIR=%~dp0"

:: Verificar si existe el archivo .env
IF NOT EXIST "%CURRENT_DIR%.env" (
    echo [ERROR] No se encuentra el archivo .env en %CURRENT_DIR%
    echo Por favor, crea un archivo .env con tu CLIMA_API_KEY
    pause
    exit /b
)

echo Iniciando Clima CLI App...

:: Abrir una nueva ventana de terminal y ejecutar Docker
:: -it: Modo interactivo
:: --rm: Eliminar contenedor al salir
:: -v: Montar el archivo .env para persistencia
start "Clima CLI 🌦️" cmd /c "docker run -it --rm -v "%CURRENT_DIR%.env":/root/.env clima-cli:2.0"

ENDLOCAL
