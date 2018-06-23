FROM golang:1-alpine AS build
LABEL maintainer="UshioShugo<ushio.s@gmail.com>"

ENV APP_PATH=${GOPATH}/src/github.com/ushios/golem

RUN apk add --no-cache --virtual .dep \
	git openssh-client && \
	go get github.com/ushios/golem/...

WORKDIR ${APP_PATH}


FROM alpine
RUN apk add --no-cache \
	ca-certificates
COPY --from=build /go/bin/golem /usr/local/bin/golem
ENTRYPOINT ["/usr/local/bin/golem"]
CMD ["-help"]
