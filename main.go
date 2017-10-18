package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	// FakeTSDB listen port
	port = 8089

	// Forward to influxDB http api
	influxAddress = "http://127.0.0.1:8086"

	cache = 3

	influxDatabase = "test"

	debug = false
)

func init() {
	flag.IntVar(&port, "port", port, "Fake proxy listen port.")
	flag.IntVar(&cache, "cache", cache, "Number of batch items send to influx.")
	flag.StringVar(&influxAddress, "influxAddr", influxAddress, "InfluxDB HTTP API address.")
	flag.StringVar(&influxDatabase, "influxDatabase", influxDatabase, "InfluxDB Database.")
	flag.BoolVar(&debug, "debug", debug, "Enable debug mode.")
	flag.Parse()

	if !strings.HasPrefix(influxAddress, "http://") && !strings.HasPrefix(influxAddress, "https://") {
		fmt.Println("ERROR: influxAddress must contain a prefix http[s]://")
		os.Exit(1)
	}

	influxRevciver = NewInfluxDBReciver(influxAddress, cache)
	influxRevciver.ListenTask()
}

func main() {
	NewTSDBServer(fmt.Sprintf("0.0.0.0:%d", port), influxAddress).TcpServer.Listen()
}
