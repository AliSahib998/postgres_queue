FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

EXPOSE 8085

CMD ["./main"]
