FROM alpine:3.6

RUN apk add --no-cache ca-certificates

ADD gc-operator /usr/local/bin/gc-operator

RUN adduser -D gc-operator
USER gc-operator
