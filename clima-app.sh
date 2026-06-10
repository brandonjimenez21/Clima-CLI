#!/bin/bash

# Obtener la ruta del directorio actual
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Verificar si existe el archivo .env
if [ ! -f "$CURRENT_DIR/.env" ]; then
    echo "[ERROR] No se encuentra el archivo .env en $CURRENT_DIR"
    echo "Por favor, crea un archivo .env con tu CLIMA_API_KEY"
    read -p "Presiona Enter para salir..."
    exit 1
fi

echo "Iniciando Clima CLI App..."

# Intentar abrir en una nueva terminal (funciona en la mayoría de distros Linux con x-terminal-emulator o gnome-terminal)
if command -v x-terminal-emulator &> /dev/null; then
    x-terminal-emulator -e "docker run -it --rm -v $CURRENT_DIR/.env:/root/.env clima-cli:2.0"
elif command -v gnome-terminal &> /dev/null; then
    gnome-terminal -- docker run -it --rm -v $CURRENT_DIR/.env:/root/.env clima-cli:2.0
else
    # Si no se encuentra un emulador, ejecutar en la terminal actual
    docker run -it --rm -v $CURRENT_DIR/.env:/root/.env clima-cli:2.0
fi
