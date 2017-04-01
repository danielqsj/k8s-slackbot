FROM alpine:3.5

RUN apk add --no-cache ca-certificates bash

COPY . /app/build

RUN    /app/build/build/build-go.sh \
    && /app/build/build/build.sh \
    && /app/build/build/finalize.sh \
    && rm -rf /app/build

ENTRYPOINT ["/app/k8s-slackbot"]