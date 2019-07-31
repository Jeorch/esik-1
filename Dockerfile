#源镜像
FROM    golang:1.12.4-alpine

# 安装git
RUN     apk add --no-cache git gcc musl-dev mercurial bash gcc g++ make pkgconfig openssl-dev

#LABEL 更改version后，本地build时LABEL以上的Steps使用Cache
LABEL   maintainer="czhang@pharbers.com" PhAuthServer.version="0.0.1"

# 设置工程配置文件的环境变量 && 开启go-module
ENV     PROJECT_NAME esik
ENV     GITHUB_URL https://github.com/Jeorch
ENV     BM_KAFKA_CONF_HOME $GOPATH/$PROJECT_NAME/resources
ENV     LOGGER_USER $HOSTNAME
ENV     LOGGER_DEBUG true
ENV     LOG_PATH $GOPATH/go.log
ENV     ESIK_TOPIC esik
ENV     ESIK_MOUNT_POINT /
ENV     ESIK_TICKER_MS 10000
ENV     GO111MODULE on
ENV     PKG_CONFIG_PATH /usr/lib/pkgconfig

# 下载rdkafka
RUN git clone https://github.com/edenhill/librdkafka.git $GOPATH/librdkafka

WORKDIR $GOPATH/librdkafka
RUN ./configure --prefix /usr  && \
make && \
make install

# 下载工程
RUN     git clone $GITHUB_URL/$PROJECT_NAME $GOPATH/$PROJECT_NAME

# 设置工作目录
WORKDIR $GOPATH/$PROJECT_NAME

# 构建可执行文件
RUN     go build -a && go install

# 设置工作目录
WORKDIR $GOPATH/bin

ENTRYPOINT ["esik"]