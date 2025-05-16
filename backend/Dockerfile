# Etapa 1: Build da aplicação usando a imagem base personalizada
FROM billbridge-base:1.0 AS builder

WORKDIR /app

# Copia os arquivos go.mod e go.sum primeiro (para usar o cache)
COPY go.mod go.sum ./

# Baixa as dependências (usará o cache se go.mod e go.sum não mudarem)
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila a aplicação (sem CGO, para compatibilidade com Alpine/scratch)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Etapa 2: Imagem final, enxuta, com o binário apenas
FROM alpine:3.21.3

# Instala certificados raiz e compatibilidade com glibc
RUN apk --no-cache add ca-certificates libc6-compat

WORKDIR /app

# Copia o binário da etapa de build
COPY --from=builder /app/main .

# Expõe a porta usada pela aplicação
EXPOSE 8080

# Comando de execução
CMD ["./main"]

