# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /Twiddit_auth_ms

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /twiddit_auth_ms

EXPOSE 1414

CMD [ "/twiddit_auth_ms" ]