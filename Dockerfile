# Usar uma imagem oficial do Go 1.22.5 como base
FROM golang:1.22.5-alpine AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o código fonte para o diretório de trabalho
COPY . .

# Baixar as dependências
RUN go mod download

# Criar o diretório /app/bin para armazenar os binários
RUN mkdir -p ./bin

# Compilar o binário principal
RUN go build -o ./bin/serve ./cmd/serve/main.go

# Usar uma imagem mínima para executar a aplicação
FROM ubuntu:22.04

# Criar diretório para a aplicação
WORKDIR /app

# Copiar os binários da fase de build para a imagem final
COPY --from=builder /app/bin/serve .

# Atualizar e instalar apenas as dependências necessárias, incluindo o Chromium
RUN apt-get update && \
    apt-get install -y \
        chromium-browser \
        libnss3 \
        libx11-xcb1 \
        libxcomposite1 \
        libxdamage1 \
        libxrandr2 \
        libxtst6 \
        libatspi2.0-0 \
        libgtk-3-0 \
        libgbm1 \
        fonts-liberation \
        xdg-utils

ENV PORT=8080

# Expor a porta que a aplicação principal irá rodar (ajuste conforme necessário)
EXPOSE 8080

# Executar o binário da aplicação
CMD ["sh", "-c", "./serve"]
