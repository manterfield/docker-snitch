FROM golang:1.8-alpine

COPY . /go/src/github.com/manterfield/docker-snitch/

RUN apk add --no-cache --virtual .build-deps \
    build-base \
    git && \
    go get github.com/manterfield/docker-snitch/... && \
    rm -rf /go/src /go/pkg && \
    apk del .build-deps

ENV DOCKER_ENDPOINT "unix:///tmp/docker.sock"

EXPOSE 8080

ENTRYPOINT ["snitch-server"]
