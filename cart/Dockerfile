FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY cart cart

RUN go build -o cart_service ./cart/cmd/

FROM alpine:3.13

WORKDIR /src

COPY --from=compiler /src/app/cart_service /src/app
CMD ["/src/app"]

