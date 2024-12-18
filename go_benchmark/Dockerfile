FROM golang:1.23.4-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor ./vendor

COPY . .

RUN go build -mod=vendor -o main .

RUN chmod +x /app/main

EXPOSE 3000

CMD ["/app/main"]