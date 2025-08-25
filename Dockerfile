# ---------- Etapa 1: Build ----------
FROM golang:1.24 AS builder

# Crear carpeta de trabajo
WORKDIR /app

# Copiar y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el resto del proyecto
COPY . .

# Compilar binario optimizado para Linux (Cloud Run)
RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# ---------- Etapa 2: Imagen final ligera ----------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copiar binario desde la etapa anterior
COPY --from=builder /app/server .

# Copiar docs de swagger si las necesitas en runtime
COPY --from=builder /app/docs ./docs

# Cloud Run inyecta $PORT
ENV PORT=8080

EXPOSE 8080

# Ejecutar el binario
CMD ["./server"]