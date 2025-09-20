FROM golang:1.24-alpine

WORKDIR /src
COPY ./ ./
RUN go mod download