FROM kingproxj/golang:2.0.0
MAINTAINER xj
WORKDIR /usr/bin
ADD . /usr/bin
RUN go build main.go
LABEL workdir="/usr/bin" exposeports="8888" cmdopts="" gitlastcommit=""
EXPOSE 8888
ENTRYPOINT ["/usr/bin/main"]
