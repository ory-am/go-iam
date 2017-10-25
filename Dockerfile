FROM golang:1.9-alpine

RUN apk add --no-cache git build-base
ENV DEP_VERSION 0.3.2
RUN curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o $GOPATH/bin/dep
RUN chmod +x $GOPATH/bin/dep

WORKDIR /go/src/github.com/ory/hydra

ADD ./Gopkg.lock ./Gopkg.lock
RUN dep ensure

ADD . .
RUN go install .

ENTRYPOINT /go/bin/hydra host

EXPOSE 4444
