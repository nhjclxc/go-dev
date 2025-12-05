FROM golang:1.21-alpine AS builder

WORKDIR /app
# COPY main.go .
COPY . .
RUN go build -o backend main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/backend .
EXPOSE 8881
CMD ["./backend"]
