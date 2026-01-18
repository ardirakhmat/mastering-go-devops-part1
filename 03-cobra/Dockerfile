FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o mycli
FROM scratch
COPY --from=builder /app/mycli .
ENTRYPOINT ["./mycli"]
