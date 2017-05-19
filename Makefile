SHELL := /bin/bash
.PHONY: docker shell

docker:
	docker build -t code-challenge:latest .

shell:
	docker run --rm -it -P --name code-challenge \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-v "`pwd`":/go/src/github.com/replicatedcom/code-challenge \
	code-challenge:latest

test:
	ginkgo -r -v