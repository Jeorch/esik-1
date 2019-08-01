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
|   ESIK_TOPIC    |   'Extract System Info to Kafka' Topic    |
|   ESIK_MOUNT_POINT    |   'Extract System Info to Kafka' Disk Mount Point    |
|   ESIK_TICKER_MS    |   'Extract System Info to Kafka' Ticker Millisecond    |

***2.Run esik.***
```cassandraql
$ go build -a
$ go run esik.go
```

***3.Run esik in docker container.***
```shell script
$ docker build -t="esik" .
$ docker run --name esik esik -v ${your_kafka_conf_path}/secrets:/go/esik/secrets:ro
```
or
```shell script
$ docker-compose -f docker-compose.yml up
```
