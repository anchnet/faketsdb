# faketsdb
Forward data to InfluxDB with fake OpenTSDB protocol in Open-Falcon

## Usage
```sh
go get github.com/51idc/faketsdb
```

## Helper
```sh
🍺 eagle [~] → faketsdb -h
Usage of faketsdb:
  -cache int
    	Number of batch items send to influx. (default 3)
  -influxAddr string
    	InfluxDB HTTP API address. (default "http://127.0.0.1:8086")
  -influxDatabase string
    	InfluxDB Database. (default "test")
  -port int
    	Fake proxy listen port. (default 8089)
 ```
