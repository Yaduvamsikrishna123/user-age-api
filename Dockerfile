FROM golang:1.25.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/server

EXPOSE 3000

CMD ["./main"]