FROM golang:1.17.0-alpine3.13 as compiler

WORKDIR /src/app

COPY go.mod go.sum ./

COPY favourites favourites

RUN go build -o favourites_service ./favourites/cmd/

FROM alpine:3.13

WORKDIR /src

COPY --from=compiler /src/app/favourites_service /src/app
CMD ["/src/app"]
