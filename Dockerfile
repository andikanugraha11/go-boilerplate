# builder image
FROM golang:1.17-alpine3.14 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o go-boiler-plate .

FROM alpine:3.14
RUN apk --no-cache --update add ca-certificates
RUN apk add --no-cache tzdata
COPY --from=builder /build/go-boiler-plate .

CMD  /go-boiler-plate