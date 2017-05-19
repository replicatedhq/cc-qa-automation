FROM golang:1.7.5

RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
RUN apt-get update
RUN apt-get install -y vim

ENV PROJECTPATH=/go/src/github.com/replicatedcom/code-challenge

ENV LOG_LEVEL DEBUG

COPY . $PROJECTPATH

WORKDIR $PROJECTPATH

CMD ["/bin/bash"]
