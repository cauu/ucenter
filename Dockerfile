FROM registry.cn-hangzhou.aliyuncs.com/cauu/dev:wf-server.0.0.1

MAINTAINER Martin "cauu@163.com"

WORKDIR $GOPATH/bin

RUN cd $GOPATH/src \
  && git clone https://github.com/cauu/ucenter.git --depth=1 \
  && cd ucenter \
  && go install

EXPOSE 80

EXPOSE 8080

ENTRYPOINT ["./ucenter"]
