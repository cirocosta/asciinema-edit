FROM golang:alpine AS builder

ADD ./ /go/src/github.com/cirocosta/asciinema-edit/
WORKDIR /go/src/github.com/cirocosta/asciinema-edit

RUN set -x && \
	apk add --update make

RUN set -x && \
        make test && \
        go build \
                -tags netgo -v -a \
                -ldflags "-X main.version=$(cat ./VERSION) -extldflags \"-static\"" && \
        mv ./asciinema-edit /usr/bin/asciinema-edit

FROM alpine
COPY --from=builder /usr/bin/asciinema-edit /usr/bin/asciinema-edit

