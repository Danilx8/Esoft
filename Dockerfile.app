FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

COPY ./app ./app
COPY .env ./app

RUN go build -o main ./app/cmd/main.go


FROM alpine

WORKDIR /build

COPY .env /build
COPY --from=builder /build/main /build/main
RUN chmod +x /build/main

CMD ["/build/main"]
