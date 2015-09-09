FROM     ubuntu:trusty
MAINTAINER Nurza nurza@90tech.fr

RUN (apt-get update && apt-get install -y -q wget git && apt-get -y -q autoclean && apt-get -y -q autoremove)

RUN (wget -O /tmp/go.tar.gz https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz)
RUN (cd /tmp && tar xf go.tar.gz && mv go /usr/local && rm /tmp/go.tar.gz)

ENV GOPATH=/go

RUN mkdir /go
RUN /usr/local/go/bin/go get "github.com/nurza/logo"
RUN /usr/local/go/bin/go get "github.com/90TechSAS/libgo-docker-guard"
RUN /usr/local/go/bin/go get "github.com/gorilla/mux"
RUN /usr/local/go/bin/go get "gopkg.in/yaml.v2"

ADD dgp /dgp

RUN (cd dgp && /usr/local/go/bin/go build -o dg-probe main.go)

RUN mkdir /var/lib/docker
VOLUME /var/run/ /var/run/
VOLUME /var/lib/docker /var/lib/docker

EXPOSE 8123

ENTRYPOINT /dgp/dg-probe -f /dgp/config.yaml
