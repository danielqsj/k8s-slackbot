FROM alpine:3.5

RUN apk add --no-cache ca-certificates bash

COPY . /app/k8s-slackbot/build
WORKDIR /app/k8s-slackbot


RUN    ./build/docker/build-go.sh \
    && ./build/docker/build.sh \
    && ./build/docker/finalize.sh

ENTRYPOINT ["/k8s-slackbot"]