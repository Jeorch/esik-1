## Esik
Esik means "Extract System Info to Kafka". Go version is 12.

## Usage
***1.Set some env.***

|   env_key |   description |
|   ----    |   ----    |
|   LOGGER_DEBUG    |   true or false    |
|   LOG_PATH    |   log store path    |
|   LOGGER_USER    |   log user    |
|   BM_KAFKA_CONF_HOME    |   Kafka conf for BM kafka lib    |
|   ESIK_TOPIC    |   'Extract System Info to Kafka' Topic    |
|   ESIK_MOUNT_POINT    |   'Extract System Info to Kafka' Disk Mount Point    |
|   ESIK_TICKER_MS    |   'Extract System Info to Kafka' Ticker Millisecond    |
|   GO111MODULE    |   Set value 'on' used to enable go module    |
|   PKG_CONFIG_PATH    |   Dependence of rdkafka    |

***2.Run esik.***