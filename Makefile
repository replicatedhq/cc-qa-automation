SHELL := /bin/bash
.PHONY: docker shell

docker:
	docker build -t code-challenge:latest .

test: docker
	docker run --rm \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-v "`pwd`":/go/src/github.com/replicatedcom/code-challenge \
	code-challenge:latest
