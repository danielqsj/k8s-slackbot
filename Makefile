all: push

TAG = 1.0
PREFIX = danielqsj/k8s-slackbot

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef VERSION
  VERSION := git-$(shell git rev-parse --short HEAD)
endif

# base package. It contains the common and backends code
PKG := "github.com/danielqsj/k8s-slackbot"

GO_LIST_FILES=$(shell go list ${PKG}/... | grep -v vendor | grep -v -e "test/e2e")

.PHONY: k8s-slackbot
k8s-slackbot: clean
	CGO_ENABLED=0 GOOS=linux GO15VENDOREXPERIMENT=1 go build -a -installsuffix cgo -ldflags \
		"-s -w -X main.version=${VERSION} -X main.gitRepo=${REPO_INFO}" \
		-o k8s-slackbot \
		./main.go

.PHONY: container
container: k8s-slackbot
	docker build -t $(PREFIX):$(TAG) .

.PHONY: push
push: container
	docker push $(PREFIX):$(TAG)

.PHONY: clean
clean:
	rm -f k8s-slackbot

.PHONY: cover
cover:
	@go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c
	gover
	goveralls -coverprofile=gover.coverprofile -service travis-ci

.PHONY: vet
vet:
	@go vet ${GO_LIST_FILES}

.PHONY: fmt
fmt:
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c

.PHONY: lint
lint:
	@go list -f '{{if len .TestGoFiles}}"golint -min_confidence=0.85 {{.Dir}}/..."{{end}}' ${GO_LIST_FILES} | xargs -L 1 sh -c