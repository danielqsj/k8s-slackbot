FROM alpine:latest

RUN apk add --no-cache ca-certificates
ADD kubernetes-slackbot /
ENTRYPOINT ["/kubernetes-slackbot"]
