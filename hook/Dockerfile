FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY hook hook

RUN go build -o hook_service ./hook/cmd/

FROM alpine:3.13

WORKDIR /src

ARG PORT="4035"
ENV PORT=$PORT
EXPOSE $PORT

COPY --from=compiler /src/app/hook_service /src/app
CMD ["/src/app"]
