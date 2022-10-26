FROM alpine:latest

FROM golang:1.17-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o bank_app ./cmd/main.go

CMD ["./bank_app"]