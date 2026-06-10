# Etapa 1: Construcción
FROM golang:1.24-alpine AS builder

# Instalar dependencias necesarias para la compilación
RUN apk add --no-cache git

WORKDIR /app

# Copiar el archivo de módulo
COPY go.mod ./

# Copiar el código fuente para que go mod tidy funcione
COPY . .

# Generar go.sum y descargar dependencias
RUN go mod tidy && go mod download

# Compilar el binario estático
RUN CGO_ENABLED=0 GOOS=linux go build -o clima-cli ./cmd/clima-cli

# Etapa 2: Imagen Final Ligera
FROM alpine:3.19

# Agregar certificados CA para peticiones HTTPS
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copiar el binario desde la etapa de construcción
COPY --from=builder /app/clima-cli .

# Ejecutar el binario
ENTRYPOINT ["./clima-cli"]
