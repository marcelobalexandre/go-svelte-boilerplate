FROM golang:1.23-alpine

ARG GO_ENV
ENV GO_ENV $GO_ENV

ARG CGO_ENABLED
ENV CGO_ENABLED $CGO_ENABLED

RUN apk update && apk add --no-cache bash build-base

WORKDIR /app

COPY go.mod go.sum .
RUN go mod tidy
RUN go install github.com/air-verse/air@latest

COPY . .
RUN make build

EXPOSE 8080

CMD ["bin/dev"]
