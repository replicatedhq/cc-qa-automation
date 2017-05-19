FROM golang:1.7

RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega

ENV PROJECTPATH=/go/src/github.com/replicatedcom/code-challenge

ENV LOG_LEVEL DEBUG

WORKDIR $PROJECTPATH

CMD ["ginkgo", "-r", "-v"]
