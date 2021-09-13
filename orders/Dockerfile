FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY orders orders

RUN go build -o orders_service ./orders/cmd/

FROM alpine:3.13

WORKDIR /src

COPY --from=compiler /src/app/orders_service /src/app
CMD ["/src/app"]
