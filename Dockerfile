FROM golang:1.15.8-alpine3.13 as builder
RUN mkdir /build
COPY . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o bitcoin-api-server .


# generate clean, final image for end users
FROM alpine:3.13
COPY --from=builder /build/bitcoin-api-server .

CMD ./bitcoin-api-server $PORT