FROM alpine:latest
ADD kubernetes-slackbot /
ENTRYPOINT ["/kubernetes-slackbot"]
