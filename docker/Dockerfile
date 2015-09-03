FROM     ubuntu:trusty
MAINTAINER Nurza nurza@90tech.fr

RUN (apt-get update && apt-get install -y -q wget git && apt-get -y -q autoclean && apt-get -y -q autoremove)

RUN (wget -O /tmp/go.tar.gz https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz)
RUN (cd /tmp && tar xvf go.tar.gz && mv go /usr/local)

RUN git clone https://github.com/90TechSAS/docker-guard-probe.git
RUN mkdir /go
RUN GOPATH=/go /usr/local/go/bin/go get "github.com/nurza/logo"
RUN GOPATH=/go /usr/local/go/bin/go get "github.com/90TechSAS/libgo-docker-guard"
RUN GOPATH=/go /usr/local/go/bin/go get "github.com/gorilla/mux"
RUN GOPATH=/go /usr/local/go/bin/go get "gopkg.in/yaml.v2"
RUN (cd docker-guard-probe && GOPATH=/go /usr/local/go/bin/go build -o dg-probe main.go)

RUN mkdir /var/lib/docker
VOLUME /var/run/ /var/run/
VOLUME /var/lib/docker /var/lib/docker

ADD config.yaml /docker-guard-probe/config.yaml

EXPOSE 8123

ENTRYPOINT /docker-guard-probe/dg-probe -f /docker-guard-probe/config.yaml
