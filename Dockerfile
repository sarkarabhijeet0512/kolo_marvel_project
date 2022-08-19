FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd/
RUN apk update && apk add curl
RUN adduser -S -D -H -h /app appuser
USER appuser