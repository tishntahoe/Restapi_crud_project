FROM golang:1.23-alpine AS builder

WORKDIR /app

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем код и собираем бинарник
COPY . .
RUN go build -o server ./cmd


FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
