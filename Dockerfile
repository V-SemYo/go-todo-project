# Сборка
FROM golang:1.25.0-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o scheduler ./main.go

# Запуск
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/scheduler .
COPY --from=builder /app/web ./web
EXPOSE 7540
CMD ["./scheduler"]