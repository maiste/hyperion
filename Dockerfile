FROM golang:1.14-alpine
LABEL maintainer="Étienne \"maiste\" Marais <etiennemarais@maiste.fr>"

WORKDIR /go/src/app
ENV GOPATH="${GOPATH}:/go/src/app/"

Run apk add git
Run go get github.com/gorilla/mux

COPY . .
RUN go build src/hyperion/main.go

RUN mkdir /templates
RUN mv /go/src/app/main /
WORKDIR /

ENTRYPOINT ["./main"]
