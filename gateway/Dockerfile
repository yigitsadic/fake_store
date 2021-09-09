FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY gateway gateway
COPY auth auth

RUN go build -o gateway_service ./gateway/cmd/

FROM alpine:3.13

WORKDIR /src

COPY --from=compiler /src/app/gateway_service /src/app
CMD ["/src/app"]
