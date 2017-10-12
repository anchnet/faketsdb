package main

import (
	"bytes"
	"math/rand"
	"net"
	"testing"
)

func TestParseTSDBItem(t *testing.T) {
	item := &TsdbItem{
		Metric: "sys.cpu.nice",
		Tags: map[string]string{
			"host":  "web01",
			"host2": "web02",
		},
		Value:     42.5,
		Timestamp: 1365465600,
	}

	bufferStr := item.TsdbString()

	t.Log(bufferStr)

	if item, err := ParseTSDBItem(bufferStr); err != nil {
		t.Error(err)
	} else {
		t.Logf("%+v", item)
	}
}

func BenchmarkPutMessage(b *testing.B) {

	conn, err := net.Dial("tcp", "127.0.0.1:8089")
	if err != nil {
		b.Fatal(err)
	}

	defer conn.Close()

	item := &TsdbItem{
		Metric: "sys.cpu.nice",
		Tags: map[string]string{
			"host":  "web01",
			"host2": "web02",
		},
		Value:     42.5,
		Timestamp: 1365465600,
	}

	for i := 0; i < b.N; i++ {
		item.Timestamp = item.Timestamp + int64(i)
		item.Value = rand.Float64()

		var tsdbBuffer bytes.Buffer
		tsdbBuffer.WriteString(item.TsdbString())
		tsdbBuffer.WriteString("\n")

		conn.Write(tsdbBuffer.Bytes())
	}

}
