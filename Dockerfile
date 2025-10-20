# Multi-stage build para otimizar o tamanho da imagem final
# Stage 1: Build stage
FROM golang:1.24-alpine AS builder

# Instalar dependências necessárias para build
RUN apk add --no-cache git ca-certificates tzdata

# Definir diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum primeiro (para cache de dependências)
COPY go.mod go.sum ./

# Download das dependências (cached se go.mod/go.sum não mudaram)
RUN go mod download

# Copiar código fonte
COPY . .

# Build da aplicação com otimizações
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main ./cmd/api

# Stage 2: Runtime stage
FROM alpine:latest

# Instalar certificados CA para HTTPS e timezone data
RUN apk --no-cache add ca-certificates tzdata

# Criar usuário não-root para segurança
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Criar diretório da aplicação
WORKDIR /app

# Copiar binário do stage anterior
COPY --from=builder /app/main .

# Mudar ownership para usuário não-root
RUN chown appuser:appgroup /app/main

# Mudar para usuário não-root
USER appuser

# Expor porta
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8000/helthz || exit 1

# Comando para executar a aplicação
CMD ["./main"]