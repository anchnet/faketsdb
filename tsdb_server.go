package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

type TsdbItem struct {
	Metric    string            `json:"metric"`
	Tags      map[string]string `json:"tags"`
	Value     float64           `json:"value"`
	Timestamp int64             `json:"timestamp"`
}

func (this *TsdbItem) String() string {
	return fmt.Sprintf(
		"<Metric:%s, Tags:%v, Value:%v, TS:%d>",
		this.Metric,
		this.Tags,
		this.Value,
		this.Timestamp,
	)
}

func (this *TsdbItem) TsdbString() (s string) {
	s = fmt.Sprintf("put %s %d %.3f ", this.Metric, this.Timestamp, this.Value)

	for k, v := range this.Tags {
		key := strings.ToLower(strings.Replace(k, " ", "_", -1))
		value := strings.Replace(v, " ", "_", -1)
		s += key + "=" + value + " "
	}

	return s
}

var ErrInvalidProtocol = errors.New("Invalid protocol")

func ParseTSDBItem(str string) (*TsdbItem, error) {
	segments := strings.Split(str, " ")
	if len(segments) < 5 {
		return nil, errors.New("Invalid protocol: number of segments less than 5.")
	}

	if strings.ToLower(segments[0]) != "put" {
		return nil, errors.New("Invalid protocol: only suppurt \"put\" operate.")
	}

	item := &TsdbItem{
		Metric: segments[1],
	}

	var err error
	item.Timestamp, err = strconv.ParseInt(segments[2], 10, 64)
	if err != nil {
		return nil, errors.New("Invalid protocol: failed to parse timestamp.")
	}

	item.Value, err = strconv.ParseFloat(segments[3], 64)
	if err != nil {
		return nil, errors.New("Invalid protocol: failed to parse value.")
	}

	item.Tags = make(map[string]string)

	for i, pair := range segments {
		if i < 4 || len(pair) == 0 {
			// Skip top 4 elements and last one
			continue
		}

		tmp := strings.Split(pair, "=")
		if len(tmp) != 2 {
			continue
		}
		item.Tags[tmp[0]] = tmp[1]
	}

	return item, nil
}

type TSDBServer struct {
	TcpServer *server
}

func NewTSDBServer(address, influxAddr string) *TSDBServer {
	tsdbServer := &TSDBServer{}

	tcpServ := NewTCPServer(address)
	tcpServ.OnNewClient(tsdbServer.OnConnected)
	tcpServ.OnNewMessage(tsdbServer.OnMessage)
	tsdbServer.TcpServer = tcpServ

	return tsdbServer
}

func (s *TSDBServer) OnConnected(c *Client) {
	//c.Send("Hello\n")
}

var sema = NewSemaphore(2000)

func (s *TSDBServer) OnMessage(c *Client, message string) {

	sema.Acquire()
	go func(message string) {
		defer sema.Release()

		item, err := ParseTSDBItem(message)
		if err != nil {
			log.Printf("ERROR: %s", err.Error())
			return
		}
		log.Printf("Message: %s\n", item)

		fields := map[string]interface{}{"value": item.Value}
		t := time.Unix(item.Timestamp, 0)

		point, err := client.NewPoint(item.Metric, item.Tags, fields, t)
		if err != nil {
			log.Printf("ERROR: failed to NewPoint %s\n", err.Error())
			return
		}
		influxDBChan <- point
	}(message)

}
