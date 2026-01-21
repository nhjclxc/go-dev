FROM docker.1ms.run/library/golang:latest

WORKDIR /app

COPY scheduler.go .

RUN go mod init scheduler && go mod tidy

CMD ["go", "run", "scheduler.go"]
