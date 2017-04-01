FROM alpine:3.5

RUN apk add --no-cache ca-certificates bash

COPY . /app

RUN    /app/build/build-go.sh \
    && /app/build/build.sh \
    && /app/build/finalize.sh \
    && rm -rf /app

ENTRYPOINT ["/k8s-slackbot"]