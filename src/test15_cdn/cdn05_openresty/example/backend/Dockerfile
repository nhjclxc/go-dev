FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY main.go /app/main.go

RUN go build -o backend main.go


FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/backend /app/backend

RUN mkdir -p /app/logs

EXPOSE 8080
ENV PORT=8080

CMD ["/app/backend"]