package main

import (
	"esik/models"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmerror"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"os"
	"strconv"
	"strings"
	"time"
)

var bkc *bmkafka.BmKafkaConfig
var t *time.Ticker
var err error

func main() {
	bmlog.StandardLogger().Info("Esik Start.")
	bkc, err = bmkafka.GetConfigInstance()
	if err != nil {
		panic(err.Error())
	}

	si := new(models.SystemInfo)

	tickerStr := os.Getenv("ESIK_TICKER_MS")
	if tickerStr == "" {
		tickerStr = "10000"	//10s
	}
	ticker, err := strconv.ParseInt(tickerStr, 10, 64)
	if err != nil {
		panic(err.Error())
	}

	t = time.NewTicker(time.Duration(ticker) * time.Millisecond)
	defer clear()
	for {
		select {
		case <-t.C:
			err = si.ExtractSystemInfo()
			if err != nil {
				panic(err.Error())
			}
			err = producceSi(*si)
			if err != nil {
				panic(err.Error())
			}
		}
	}

}

func producceSi(si models.SystemInfo) (err error) {

	fmt.Println(si)

	topic := os.Getenv("ESIK_TOPIC")
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		bmlog.StandardLogger().Error("no HOSTNAME env set.")
	}
	if topic == "" {
		topic = "esik_" + strings.ReplaceAll(hostname, ".", "_")
	}

	var rawMetricsSchema = fmt.Sprint(`{"type": "record","name": "`, topic, `","fields": `,
		`[{"name": "time", "type": "string"},{"name": "hostname",  "type": "string" },{"name": "ip",  "type": "string" },`,
		`{"name": "cpu",  "type": "string" },{"name": "memory",  "type": "string" },{"name": "disk",  "type": "string" },{"name": "receive",  "type": "string" },{"name": "transmit",  "type": "string" }]}`)
	//var rawMetricsSchema = `{"type": "record","name": "esik","fields": [{"name": "time", "type": "string"},{"name": "hostname",  "type": "string" },{"name": "ip",  "type": "string" },
	//		{"name": "cpu",  "type": "string" },{"name": "memory",  "type": "string" },{"name": "disk",  "type": "string" },{"name": "receive",  "type": "string" },{"name": "transmit",  "type": "string" }]}`

	encoder := kafkaAvro.NewKafkaAvroEncoder(bkc.SchemaRepositoryUrl)
	schema, err := avro.ParseSchema(rawMetricsSchema)
	bmerror.PanicError(err)
	record := avro.NewGenericRecord(schema)
	bmerror.PanicError(err)
	record.Set("time", si.Time)
	record.Set("hostname", si.Hostname)
	record.Set("ip", si.Ip)
	record.Set("cpu", models.PercentFormat(si.CpuUsage))
	record.Set("memory", models.PercentFormat(si.MemUsage))
	record.Set("disk", models.PercentFormat(si.DiskUsage))
	record.Set("receive", models.ByteFormat(float64(si.NetStatus.Receive)))
	record.Set("transmit", models.ByteFormat(float64(si.NetStatus.Transmit)))
	recordByteArr, err := encoder.Encode(record)
	bmerror.PanicError(err)

	if err != nil {
		panic(err.Error())
	}

	bkc.Produce(&topic, recordByteArr)

	return
}

func clear() {
	bkc = nil
	if t != nil {
		t.Stop()
	}
	t = nil
	fmt.Println("Esik End.")
}
