FROM golang:1.15-alpine3.12
MAINTAINER xj
WORKDIR /usr/bin
ADD . /usr/bin
RUN go build main.go
LABEL workdir="/usr/bin" exposeports="8888" cmdopts="" gitlastcommit=""
EXPOSE 8888
ENTRYPOINT ["/usr/bin/main"]
