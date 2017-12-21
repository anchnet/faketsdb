package main

import (
	"log"

	"github.com/influxdata/influxdb/client/v2"
)

var influxDBChan = make(chan *client.Point, 10)
var influxRevciver *InfluxDBReciver

type InfluxDBReciver struct {
	address string
	cache   int
}

func NewInfluxDBReciver(addr string, cache int) *InfluxDBReciver {
	return &InfluxDBReciver{address: addr, cache: cache}
}

func (r *InfluxDBReciver) AddBatchPoint(ps []*client.Point) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: influxDatabase,
	})
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoints(ps)

	client, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: r.address,
	})
	if err != nil {
		log.Fatalf("Faile to NewHTTPClient: %s", err.Error())
	}
	defer client.Close()

	return client.Write(bp)
}

func (r *InfluxDBReciver) ListenTask() {
	go func() {
		var (
			n      = 0
			points = make([]*client.Point, r.cache)
		)
		for {
			points[n] = <-influxDBChan
			if n == r.cache-1 {
				if err := influxRevciver.AddBatchPoint(points); err != nil {
					log.Printf("Failed AddBatchPoint: %s", err.Error())
				}
				n = 0
				continue
			}
			n++
		}
	}()
}
