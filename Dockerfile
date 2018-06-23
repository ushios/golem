FROM golang:1-alpine AS build
LABEL maintainer="UshioShugo<ushio.s@gmail.com>"
LABEL name="golem"
LABEL version="0.0.1"

ENV APP_PATH=${GOPATH}/src/github.com/ushios/golem

COPY . ${APP_PATH}
WORKDIR ${APP_PATH}

RUN go install ./...

FROM alpine
COPY --from=build /go/bin/golem /usr/local/bin/golem
ENTRYPOINT ["/usr/local/bin/golem"]
CMD ["-help"]
