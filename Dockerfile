FROM golang:1.14

ENV GOPATH "/tmp/go"
RUN mkdir -p $GOPATH

RUN go get github.com/kyoh86/richgo
ENV PATH="${PATH}:${GOPATH}/bin"

CMD ["tail", "-f", "/dev/null"]
