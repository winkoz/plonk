FROM golang:1.14

ENV GOPATH "/tmp/go"

CMD ["tail", "-f", "/dev/null"]
