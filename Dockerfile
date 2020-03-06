FROM golang:1.13-alpine as builder

WORKDIR /build

ENV GO111MODULE=on

RUN apk add --update git --update make

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN make build

FROM alpine

COPY --from=builder /build/ciqube /usr/bin/ciqube