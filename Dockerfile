FROM golang:1.15.6-alpine3.12 AS builder
ADD . /go/src/mtrcgo
WORKDIR /go/src/mtrcgo
RUN go get -d
RUN go build .

FROM alpine:3.12.2
WORKDIR /app
COPY --from=builder ["/go/src/mtrcgo/mtrcgo", "/go/src/mtrcgo/config.yml", "./"]
RUN ls -la

CMD ["/app/mtrcgo"]
