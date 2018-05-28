FROM golang:latest


WORKDIR $GOPATH/src/servertcp
ADD . $GOPATH/src/servertcp
RUN go get github.com/hprose/hprose-golang/rpc && \
    go get github.com/gin-gonic/gin

RUN go build .

EXPOSE 8000
EXPOSE 8888


ENTRYPOINT ["./servertcp"]