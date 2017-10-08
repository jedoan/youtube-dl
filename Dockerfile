FROM golang:1.9
ADD . /go/src/github.com/jedoan/youtube-dl
ENV token=""
WORKDIR /go/src/github.com/jedoan/youtube-dl
RUN go install
RUN go build
CMD /go/src/github.com/jedoan/youtube-dl/youtube-dl
