FROM golang:1.23.9

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN apt-get update && \
    apt-get install -y awscli && \
    rm -rf /var/lib/apt/lists/*

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main main.go
RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]