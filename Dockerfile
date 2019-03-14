FROM registry.icp.com:5000/service/devops/runtime/golang:4.5.0
MAINTAINER xj
WORKDIR /usr/bin
RUN mkdir -p /etc/config/
ARG  jar_file
ADD $jar_file /usr/bin
RUN chmod -R a+rwx /usr/bin/main
ENV TZ Asia/Shanghai
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone
LABEL workdir="/usr/bin" exposeports="8888" cmdopts="" gitlastcommit=""
EXPOSE 8888
ENTRYPOINT ["/usr/bin/main"]
