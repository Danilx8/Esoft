FROM golang:1.23-alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .

COPY ./app ./app
COPY ./configs /build

RUN go build -o main ./app/cmd/main.go


FROM alpine

WORKDIR /build

COPY ./configs /build/configs
COPY --from=builder /build/main /build/main
RUN chmod +x /build/main

CMD ["/build/main"]
