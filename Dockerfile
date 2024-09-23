FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN mkdir -p ./bin

RUN go build -o ./bin/serve ./cmd/serve/main.go

FROM browserless/chrome:latest

WORKDIR /app

COPY --from=builder /app/bin/serve .

ENV PORT=8080

EXPOSE 8080

CMD ["sh", "-c", "./serve"]
