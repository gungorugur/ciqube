FROM golang:1.13 as builder

ENV GO111MODULE=on

WORKDIR /app

RUN git clone https://github.com/gungorugur/ciqube.git

WORKDIR /app/ciqube

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ciqube cmd/ciqube.go

FROM alpine:latest

COPY --from=builder /app/ciqube/ciqube /usr/local/bin/ciqube