FROM golang:latest as BUILD
RUN apt-get update

WORKDIR /go/src/github.com/shawnritchie/interview-accountapi-master
COPY . /go/src/github.com/shawnritchie/interview-accountapi-master/.

RUN go get github.com/cucumber/godog/cmd/godog
RUN go mod download

CMD ["sh", "gotest.sh"]