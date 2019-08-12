## Esik
Esik means "Extract System Info to Kafka". Go version is 12.

## Usage
***1.Set some env.***

###### Highly dependent env.

|   env_key |   description |
|   ----    |   ----    |
|   GO111MODULE    |   Set value 'on' used to enable go module    |
|   PKG_CONFIG_PATH    |   Dependence of rdkafka    |
|   BM_KAFKA_CONF_HOME    |   Kafka conf for BM kafka lib    |

###### Low dependent env.

|   env_key |   description |
|   ----    |   ----    |
|   LOGGER_DEBUG    |   true or false    |
|   LOG_PATH    |   log store path    |
|   LOGGER_USER    |   log user    |
|   ESIK_TOPIC    |   ESIK's Topic    |
|   ESIK_MOUNT_POINT    |   ESIK's Disk Mount Point    |
|   ESIK_TICKER_MS    |   ESIK's Ticker Millisecond    |
|   ESIK_CPU_DURATION_MS    |   ESIK's CPU Duration Millisecond    |
|   ESIK_NET_DURATION_MS    |   ESIK's Net Duration Millisecond    |

***2.Install rdkafka***
```shell script
$ git clone https://github.com/edenhill/librdkafka.git $GOPATH/librdkafka

$ cd $GOPATH/librdkafka
$ ./configure --prefix /usr  && \
make && \
make install
```

***3.Run esik.***
```shell script
$ go build -a
$ go run esik.go
```
