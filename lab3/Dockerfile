
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server ./backend/cmd/main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/server ./server
# COPY ./backend/cmd/boot.env ./backend/cmd/boot.env
COPY ./frontend ./frontend
# COPY ./backend/internal/config ./backend/internal/config
# COPY ./backend/internal/templates ./backend/internal/templates
EXPOSE 8000
ENV GIN_MODE=release
CMD ["./server"]
