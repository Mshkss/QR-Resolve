# Dockerfile
FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]
