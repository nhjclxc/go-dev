FROM docker.1ms.run/library/golang:latest

WORKDIR /app

COPY upstream.go .

RUN go mod init upstream && go mod tidy

CMD ["go", "run", "upstream.go"]
