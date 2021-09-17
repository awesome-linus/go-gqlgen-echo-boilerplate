FROM golang:1.13.0-alpine3.10 as build

WORKDIR /go/app

COPY . .

RUN go mod download

RUN set -x && \
  apk update && \
  apk add --no-cache git && \
  go build -o go-graphql-echo-boilerplate

FROM alpine:3.10

WORKDIR /app

COPY --from=build /go/app/go-graphql-echo-boilerplate .

RUN set -x && \
  addgroup go && \
  adduser -D -G go go && \
  chown -R go:go /app/go-graphql-echo-boilerplate

CMD ["./go-graphql-echo-boilerplate"]
