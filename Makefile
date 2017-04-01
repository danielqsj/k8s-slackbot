all: push

TAG = 1.0
PREFIX = danielqsj/k8s-slackbot

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef VERSION
  VERSION := git-$(shell git rev-parse --short HEAD)
endif

k8s-slackbot: clean
	CGO_ENABLED=0 GOOS=linux GO15VENDOREXPERIMENT=1 go build -a -installsuffix cgo -ldflags \
		"-s -w -X main.version=${VERSION} -X main.gitRepo=${REPO_INFO}" \
		-o k8s-slackbot \
		./main.go

container: k8s-slackbot
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

clean:
	rm -f k8s-slackbot
