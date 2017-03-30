all: push

TAG = 0.1.0
PREFIX = danielqsj/kubernetes-slackbot

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef VERSION
  VERSION := git-$(shell git rev-parse --short HEAD)
endif

kubernetes-slackbot: clean
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags \
		"-s -w -X main.version=${VERSION} -X main.gitRepo=${REPO_INFO}" \
		-o kubernetes-slackbot \
		./main.go

container: kubernetes-slackbot
	docker build --no-cache -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

clean:
	rm -f kubernetes-slackbot
