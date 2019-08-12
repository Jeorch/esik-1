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
		tickerStr = "1000"	//1s
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
			fmt.Println(si)
			err = produceSi(*si)
			if err != nil {
				panic(err.Error())
			}
		}
	}

}

func produceSi(si models.SystemInfo) (err error) {

	topic := os.Getenv("ESIK_TOPIC")
	if topic == "" {
		hostname := os.Getenv("HOSTNAME")
		if hostname == "" {
			bmlog.StandardLogger().Error("no HOSTNAME env set.")
		}
		topic = "esik00" + strings.ReplaceAll(strings.ReplaceAll(hostname, ".", ""), "_", "")
	}

	var rawMetricsSchema = fmt.Sprint(`{"type": "record","name": "`, topic, `","fields": `,
		`[{"name": "time", "type": "string"},{"name": "hostname",  "type": "string" },{"name": "ip",  "type": "string" },`,
		`{"name": "cpu",  "type": "string" },{"name": "memory",  "type": "string" },{"name": "disk",  "type": "string" },{"name": "receive",  "type": "string" },{"name": "transmit",  "type": "string" }]}`)

	encoder := kafkaAvro.NewKafkaAvroEncoder(bkc.SchemaRepositoryUrl)
	schema, err := avro.ParseSchema(rawMetricsSchema)
	bmerror.PanicError(err)
	record := avro.NewGenericRecord(schema)
	bmerror.PanicError(err)
	record.Set("time", si.Time)
	record.Set("hostname", si.Hostname)
	record.Set("ip", si.Ip)
	record.Set("cpu", models.FloatFormat(si.CpuUsage))
	record.Set("memory", models.FloatFormat(si.MemUsage))
	record.Set("disk", models.FloatFormat(si.DiskUsage))
	record.Set("receive", models.FloatFormat(si.NetSpeed.Receive/1024))	//默认速度以kb/s为单位
	record.Set("transmit", models.FloatFormat(si.NetSpeed.Transmit/1024))	//默认速度以kb/s为单位
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
