FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY payment_provider payment_provider

RUN go build -o payment_provider_binary ./payment_provider/cmd/

FROM alpine:3.13

WORKDIR /src

ARG PORT="5055"
ARG BASE_URL="http://localhost:$PORT"

ENV PORT=$PORT
ENV BASE_URL=$BASE_URL

COPY --from=compiler /src/app/payment_provider_binary /src/app

EXPOSE $PORT

CMD ["/src/app"]
