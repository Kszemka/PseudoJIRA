FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY ./certs/cert.pem /app
COPY ./certs/key.pem /app

RUN go build -o main .

EXPOSE 8443

CMD ["./main"]
