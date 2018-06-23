FROM golang:1-alpine
LABEL maintainer="UshioShugo<ushio.s@gmail.com>"
LABEL name="build-golem"
LABEL version="0.0.1"

ENV APP_PATH=${GOPATH}/src/github.com/ushios/golem

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

RUN go install ./...

ENTRYPOINT ["/go/bin/golem"]
CMD ["-help"]
