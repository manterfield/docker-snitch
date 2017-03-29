FROM alpine:3.5

COPY . /go/src/github.com/manterfield/docker-snitch/

ENV GOPATH /go
ENV PATH $PATH:/go/bin:/usr/local/go/bin

RUN apk add --no-cache --virtual .build-deps \
    build-base \
    git \
    go && \
    go get github.com/manterfield/docker-snitch/... && \
    rm -rf /go/src /go/pkg && \
    apk del .build-deps

ENV DOCKER_ENDPOINT "unix:///tmp/docker.sock"

EXPOSE 8080

ENTRYPOINT ["snitch-server"]
