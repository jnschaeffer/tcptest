FROM golang:1.12 AS builder

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -o tcptest .

FROM alpine:3.10.2

COPY --from=builder /app/tcptest /app/tcptest

CMD /app/tcptest